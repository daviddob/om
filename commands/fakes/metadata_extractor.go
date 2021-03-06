// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/extractor"
)

type MetadataExtractor struct {
	ExtractMetadataStub        func(string) (extractor.Metadata, error)
	extractMetadataMutex       sync.RWMutex
	extractMetadataArgsForCall []struct {
		arg1 string
	}
	extractMetadataReturns struct {
		result1 extractor.Metadata
		result2 error
	}
	extractMetadataReturnsOnCall map[int]struct {
		result1 extractor.Metadata
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetadataExtractor) ExtractMetadata(arg1 string) (extractor.Metadata, error) {
	fake.extractMetadataMutex.Lock()
	ret, specificReturn := fake.extractMetadataReturnsOnCall[len(fake.extractMetadataArgsForCall)]
	fake.extractMetadataArgsForCall = append(fake.extractMetadataArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ExtractMetadata", []interface{}{arg1})
	fake.extractMetadataMutex.Unlock()
	if fake.ExtractMetadataStub != nil {
		return fake.ExtractMetadataStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.extractMetadataReturns.result1, fake.extractMetadataReturns.result2
}

func (fake *MetadataExtractor) ExtractMetadataCallCount() int {
	fake.extractMetadataMutex.RLock()
	defer fake.extractMetadataMutex.RUnlock()
	return len(fake.extractMetadataArgsForCall)
}

func (fake *MetadataExtractor) ExtractMetadataArgsForCall(i int) string {
	fake.extractMetadataMutex.RLock()
	defer fake.extractMetadataMutex.RUnlock()
	return fake.extractMetadataArgsForCall[i].arg1
}

func (fake *MetadataExtractor) ExtractMetadataReturns(result1 extractor.Metadata, result2 error) {
	fake.ExtractMetadataStub = nil
	fake.extractMetadataReturns = struct {
		result1 extractor.Metadata
		result2 error
	}{result1, result2}
}

func (fake *MetadataExtractor) ExtractMetadataReturnsOnCall(i int, result1 extractor.Metadata, result2 error) {
	fake.ExtractMetadataStub = nil
	if fake.extractMetadataReturnsOnCall == nil {
		fake.extractMetadataReturnsOnCall = make(map[int]struct {
			result1 extractor.Metadata
			result2 error
		})
	}
	fake.extractMetadataReturnsOnCall[i] = struct {
		result1 extractor.Metadata
		result2 error
	}{result1, result2}
}

func (fake *MetadataExtractor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.extractMetadataMutex.RLock()
	defer fake.extractMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetadataExtractor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
