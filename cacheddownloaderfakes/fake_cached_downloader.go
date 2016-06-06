// This file was generated by counterfeiter
package cacheddownloaderfakes

import (
	"io"
	"net/url"
	"sync"

	"github.com/cloudfoundry-incubator/cacheddownloader"
)

type FakeCachedDownloader struct {
	FetchStub        func(urlToFetch *url.URL, cacheKey string, checksum cacheddownloader.ChecksumInfoType, cancelChan <-chan struct{}) (stream io.ReadCloser, size int64, err error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		urlToFetch *url.URL
		cacheKey   string
		checksum   cacheddownloader.ChecksumInfoType
		cancelChan <-chan struct{}
	}
	fetchReturns struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}
	FetchAsDirectoryStub        func(urlToFetch *url.URL, cacheKey string, checksum cacheddownloader.ChecksumInfoType, cancelChan <-chan struct{}) (dirPath string, size int64, err error)
	fetchAsDirectoryMutex       sync.RWMutex
	fetchAsDirectoryArgsForCall []struct {
		urlToFetch *url.URL
		cacheKey   string
		checksum   cacheddownloader.ChecksumInfoType
		cancelChan <-chan struct{}
	}
	fetchAsDirectoryReturns struct {
		result1 string
		result2 int64
		result3 error
	}
	CloseDirectoryStub        func(cacheKey, directoryPath string) error
	closeDirectoryMutex       sync.RWMutex
	closeDirectoryArgsForCall []struct {
		cacheKey      string
		directoryPath string
	}
	closeDirectoryReturns struct {
		result1 error
	}
}

func (fake *FakeCachedDownloader) Fetch(urlToFetch *url.URL, cacheKey string, checksum cacheddownloader.ChecksumInfoType, cancelChan <-chan struct{}) (stream io.ReadCloser, size int64, err error) {
	fake.fetchMutex.Lock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		urlToFetch *url.URL
		cacheKey   string
		checksum   cacheddownloader.ChecksumInfoType
		cancelChan <-chan struct{}
	}{urlToFetch, cacheKey, checksum, cancelChan})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(urlToFetch, cacheKey, checksum, cancelChan)
	} else {
		return fake.fetchReturns.result1, fake.fetchReturns.result2, fake.fetchReturns.result3
	}
}

func (fake *FakeCachedDownloader) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeCachedDownloader) FetchArgsForCall(i int) (*url.URL, string, cacheddownloader.ChecksumInfoType, <-chan struct{}) {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].urlToFetch, fake.fetchArgsForCall[i].cacheKey, fake.fetchArgsForCall[i].checksum, fake.fetchArgsForCall[i].cancelChan
}

func (fake *FakeCachedDownloader) FetchReturns(result1 io.ReadCloser, result2 int64, result3 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCachedDownloader) FetchAsDirectory(urlToFetch *url.URL, cacheKey string, checksum cacheddownloader.ChecksumInfoType, cancelChan <-chan struct{}) (dirPath string, size int64, err error) {
	fake.fetchAsDirectoryMutex.Lock()
	fake.fetchAsDirectoryArgsForCall = append(fake.fetchAsDirectoryArgsForCall, struct {
		urlToFetch *url.URL
		cacheKey   string
		checksum   cacheddownloader.ChecksumInfoType
		cancelChan <-chan struct{}
	}{urlToFetch, cacheKey, checksum, cancelChan})
	fake.fetchAsDirectoryMutex.Unlock()
	if fake.FetchAsDirectoryStub != nil {
		return fake.FetchAsDirectoryStub(urlToFetch, cacheKey, checksum, cancelChan)
	} else {
		return fake.fetchAsDirectoryReturns.result1, fake.fetchAsDirectoryReturns.result2, fake.fetchAsDirectoryReturns.result3
	}
}

func (fake *FakeCachedDownloader) FetchAsDirectoryCallCount() int {
	fake.fetchAsDirectoryMutex.RLock()
	defer fake.fetchAsDirectoryMutex.RUnlock()
	return len(fake.fetchAsDirectoryArgsForCall)
}

func (fake *FakeCachedDownloader) FetchAsDirectoryArgsForCall(i int) (*url.URL, string, cacheddownloader.ChecksumInfoType, <-chan struct{}) {
	fake.fetchAsDirectoryMutex.RLock()
	defer fake.fetchAsDirectoryMutex.RUnlock()
	return fake.fetchAsDirectoryArgsForCall[i].urlToFetch, fake.fetchAsDirectoryArgsForCall[i].cacheKey, fake.fetchAsDirectoryArgsForCall[i].checksum, fake.fetchAsDirectoryArgsForCall[i].cancelChan
}

func (fake *FakeCachedDownloader) FetchAsDirectoryReturns(result1 string, result2 int64, result3 error) {
	fake.FetchAsDirectoryStub = nil
	fake.fetchAsDirectoryReturns = struct {
		result1 string
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCachedDownloader) CloseDirectory(cacheKey string, directoryPath string) error {
	fake.closeDirectoryMutex.Lock()
	fake.closeDirectoryArgsForCall = append(fake.closeDirectoryArgsForCall, struct {
		cacheKey      string
		directoryPath string
	}{cacheKey, directoryPath})
	fake.closeDirectoryMutex.Unlock()
	if fake.CloseDirectoryStub != nil {
		return fake.CloseDirectoryStub(cacheKey, directoryPath)
	} else {
		return fake.closeDirectoryReturns.result1
	}
}

func (fake *FakeCachedDownloader) CloseDirectoryCallCount() int {
	fake.closeDirectoryMutex.RLock()
	defer fake.closeDirectoryMutex.RUnlock()
	return len(fake.closeDirectoryArgsForCall)
}

func (fake *FakeCachedDownloader) CloseDirectoryArgsForCall(i int) (string, string) {
	fake.closeDirectoryMutex.RLock()
	defer fake.closeDirectoryMutex.RUnlock()
	return fake.closeDirectoryArgsForCall[i].cacheKey, fake.closeDirectoryArgsForCall[i].directoryPath
}

func (fake *FakeCachedDownloader) CloseDirectoryReturns(result1 error) {
	fake.CloseDirectoryStub = nil
	fake.closeDirectoryReturns = struct {
		result1 error
	}{result1}
}

var _ cacheddownloader.CachedDownloader = new(FakeCachedDownloader)