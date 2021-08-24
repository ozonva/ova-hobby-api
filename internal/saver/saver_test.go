package saver_test

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/ozonva/ova-hobby-api/internal/mocks"
	"github.com/ozonva/ova-hobby-api/internal/saver"
	"github.com/ozonva/ova-hobby-api/pkg/models"
)

var _ = Describe("Saver", func() {
	var (
		mockCtrl    *gomock.Controller
		mockFlusher *mocks.MockFlusher
		testHobby   models.Hobby
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(mockCtrl)

		testHobby = models.NewHobby("swimming", 56, models.Sports)
	})

	Context("Saver creation fails", func() {
		It("capacity must be positive number", func() {
			_, err := saver.NewSaver(0, mockFlusher, 10*time.Millisecond)
			Expect(err).To(Equal(errors.New("capacity must be positive number")))
		})
		It("flusher must be initialized", func() {
			_, err := saver.NewSaver(1, nil, 10*time.Millisecond)
			Expect(err).To(Equal(errors.New("flusher must be initialized")))
		})
	})

	Context("Calling methods after Close", func() {
		It("calling Close", func() {
			testSaver, err := saver.NewSaver(1, mockFlusher, 10*time.Millisecond)
			if err != nil {
				Fail(fmt.Sprintf("error while Saver creation: %v", err))
			}
			_, err = testSaver.Close()
			Expect(err).To(BeNil())

			_, err = testSaver.Close()
			Expect(err).To(Equal(errors.New("saver is already closed")))
		})
		It("calling Save", func() {
			testSaver, err := saver.NewSaver(1, mockFlusher, 10*time.Millisecond)
			if err != nil {
				Fail(fmt.Sprintf("error while Saver creation: %v", err))
			}
			_, err = testSaver.Close()
			Expect(err).To(BeNil())

			err = testSaver.Save(testHobby)
			Expect(err).To(Equal(errors.New("saver is already closed")))
		})
	})

	Context("saver flush hobbies", func() {
		It("non Flush calls in a goroutine", func() {
			flushFrequency := 200 * time.Millisecond
			saveCallsFrequency := 10 * time.Millisecond
			var saverCapacity uint = 2
			saveCallsQuantity := 10
			minFlushCalls := 5

			mockFlusher.EXPECT().Flush(gomock.Any()).MinTimes(minFlushCalls)
			testSaver, err := saver.NewSaver(saverCapacity, mockFlusher, flushFrequency)
			if err != nil {
				Fail(fmt.Sprintf("error while Saver creation: %v", err))
			}

			wg := sync.WaitGroup{}
			wg.Add(1)

			go func() {
				callSendTicker := time.NewTicker(saveCallsFrequency)
				defer callSendTicker.Stop()
				saveCallsCounter := 0

				for {
					<-callSendTicker.C
					Expect(testSaver.Save(testHobby)).Should(BeNil())
					saveCallsCounter += 1
					if saveCallsCounter > saveCallsQuantity {
						wg.Done()
						return
					}
				}
			}()

			wg.Wait()
		})

		It("Flush calls in a goroutine only", func() {
			flushFrequency := 10 * time.Millisecond
			saveCallsFrequency := 20 * time.Millisecond
			var saverCapacity uint = 15
			saveCallsQuantity := 10
			minFlushCalls := 20

			mockFlusher.EXPECT().Flush(gomock.Any()).MinTimes(minFlushCalls)
			testSaver, err := saver.NewSaver(saverCapacity, mockFlusher, flushFrequency)
			if err != nil {
				Fail(fmt.Sprintf("error while Saver creation: %v", err))
			}

			wg := sync.WaitGroup{}
			wg.Add(1)

			go func() {
				callSendTicker := time.NewTicker(saveCallsFrequency)
				defer callSendTicker.Stop()
				saveCallsCounter := 0

				for {
					<-callSendTicker.C
					Expect(testSaver.Save(testHobby)).Should(BeNil())
					saveCallsCounter += 1
					if saveCallsCounter > saveCallsQuantity {
						wg.Done()
						return
					}
				}
			}()

			wg.Wait()
		})
	})

})
