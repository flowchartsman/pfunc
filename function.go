//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

//
// This file borrows some implementations from
// {@link https://github.com/aws/aws-lambda-go/blob/master/lambda/handler.go}
//  - errorHandler
//  - validateArguments
//  - validateReturns
//  - NewFunction
//  - Process
//

package pfunc

import (
	"context"
	"fmt"
	"reflect"

	log "andy.dev/pfunc/logutil"
)

type function interface {
	process(ctx context.Context, input []byte) ([]byte, error)
}

type pulsarFunction func(ctx context.Context, input []byte) ([]byte, error)

func (function pulsarFunction) process(ctx context.Context, input []byte) ([]byte, error) {
	output, err := function(ctx, input)
	if err != nil {
		log.Errorf("process function error:[%s]\n", err.Error())
		return nil, err
	}

	return output, nil
}

func errorHandler(e error) pulsarFunction {
	return func(ctx context.Context, input []byte) ([]byte, error) {
		return nil, e
	}
}

func validateArguments(handler reflect.Type) (bool, error) {
	handlerTakesContext := false
	if handler.NumIn() > 2 {
		return false, fmt.Errorf("functions may not take more than two arguments, but function takes %d", handler.NumIn())
	} else if handler.NumIn() > 0 {
		contextType := reflect.TypeOf((*context.Context)(nil)).Elem()
		argumentType := handler.In(0)
		handlerTakesContext = argumentType.Implements(contextType)
		if handler.NumIn() > 1 && !handlerTakesContext {
			return false, fmt.Errorf("function takes two arguments, but the first is not Context. got %s", argumentType.Kind())
		}
	}

	return handlerTakesContext, nil
}

func validateReturns(handler reflect.Type) error {
	errorType := reflect.TypeOf((*error)(nil)).Elem()

	switch {
	case handler.NumOut() > 2:
		return fmt.Errorf("function may not return more than two values")
	case handler.NumOut() > 1:
		if !handler.Out(1).Implements(errorType) {
			return fmt.Errorf("function returns two values, but the second does not implement error")
		}
	case handler.NumOut() == 1:
		if !handler.Out(0).Implements(errorType) {
			return fmt.Errorf("function returns a single value, but it does not implement error")
		}
	}

	return nil
}

func newFunction(inputFunc any) function {
	if inputFunc == nil {
		return errorHandler(fmt.Errorf("function is nil"))
	}
	handler := reflect.ValueOf(inputFunc)
	handlerType := reflect.TypeOf(inputFunc)
	if handlerType.Kind() != reflect.Func {
		return errorHandler(fmt.Errorf("function kind %s is not %s", handlerType.Kind(), reflect.Func))
	}

	takesContext, err := validateArguments(handlerType)
	if err != nil {
		return errorHandler(err)
	}

	if err := validateReturns(handlerType); err != nil {
		return errorHandler(err)
	}

	return pulsarFunction(func(ctx context.Context, input []byte) ([]byte, error) {
		// construct arguments
		var args []reflect.Value
		if takesContext {
			args = append(args, reflect.ValueOf(ctx))
		}

		if (handlerType.NumIn() == 1 && !takesContext) || handlerType.NumIn() == 2 {
			args = append(args, reflect.ValueOf(input))
		}
		response := handler.Call(args)

		// convert return values into ([]byte, error)
		var err error
		if len(response) > 0 {
			if errVal, ok := response[len(response)-1].Interface().(error); ok {
				err = errVal
			}
		}

		var val []byte
		if len(response) > 1 {
			val = response[0].Bytes()
		}

		return val, err
	})
}

func Start[P Processor](function P) {
	fn := wrapFunction(function)
	goInstance := newGoInstance()
	err := goInstance.startFunction(fn)
	if err != nil {
		log.Fatal(fmt.Errorf("startFunction failed: %v", err))
	}
}

// GetUserConfMap provides a means to access the pulsar function's user config
// map before initializing the pulsar function
func GetUserConfMap() map[string]any {
	return NewFuncContext().userConfigs
}

// GetUserConfValue provides access to a user configuration value before
// initializing the pulsar function
func GetUserConfValue(key string) any {
	return NewFuncContext().userConfigs[key]
}
