package flusher_test

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonva/ova-hobby-api/internal/flusher"
	"github.com/ozonva/ova-hobby-api/internal/mocks"
	"github.com/ozonva/ova-hobby-api/internal/repo"
	"github.com/ozonva/ova-hobby-api/pkg/models"
)

var _ = Describe("Flusher", func() {
	var (
		mockCtrl    *gomock.Controller
		testFlusher flusher.Flusher
		repoMock    repo.Repo

		nonInitialized, singleHobby, threeHobbies []models.Hobby
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repoMock = mocks.NewMockRepo(mockCtrl)
		testFlusher = flusher.NewFlusher(2, repoMock)

		nonInitialized = nil
		singleHobby = []models.Hobby{models.NewHobby("swimming", 56, models.Sports)}
		threeHobbies = []models.Hobby{
			models.NewHobby("swimming", 56, models.Sports),
			models.NewHobby("SUP", 999, models.Sports),
			models.NewHobby("hiking", 0, models.Outdoors),
		}
	})

	Describe("Unable to Flush hobbies", func() {
		It("Non initialized slice", func() {
			Expect(testFlusher.Flush(nonInitialized)).To(Equal(nonInitialized))
		})

		It("chunkSize is bigger than len(hobbies)", func() {
			Expect(testFlusher.Flush(singleHobby)).To(Equal(singleHobby))
		})

		It("chunkSize equals zero", func() {
			zeroChunkFlusher := flusher.NewFlusher(0, repoMock)

			Expect(zeroChunkFlusher.Flush(threeHobbies)).To(Equal(threeHobbies))
		})
	})

	Describe("Successfully flushed hobbies returns nil", func() {
		var repoMockWithAddHobbiesExpected *mocks.MockRepo

		BeforeEach(func() {
			repoMockWithAddHobbiesExpected = mocks.NewMockRepo(mockCtrl)
		})

		It("Input slice is size of 3", func() {
			repoMockWithAddHobbiesExpected.EXPECT().AddHobbies(gomock.Any()).Times(2)
			testFlusher := flusher.NewFlusher(2, repoMockWithAddHobbiesExpected)

			Expect(testFlusher.Flush(threeHobbies)).To(Equal(nonInitialized))
		})

		It("Input slice is size of 1", func() {
			repoMockWithAddHobbiesExpected.EXPECT().AddHobbies(gomock.Any()).Times(1)
			oneChunkFlusher := flusher.NewFlusher(1, repoMockWithAddHobbiesExpected)

			Expect(oneChunkFlusher.Flush(singleHobby)).Should(Equal(nonInitialized))
		})
	})

})
