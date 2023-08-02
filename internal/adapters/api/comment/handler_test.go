package comment

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"someBlog/internal/domain/comment"
	mock_comment "someBlog/internal/domain/comment/mocks"
	"strconv"
	"testing"
	"time"
)

func TestHandler_Get(t *testing.T) {
	type mockBehavior func(s *mock_comment.MockService, idCom int)

	TimeTest, _ := time.Parse(time.RFC3339Nano, "2023-08-03T07:27:50.35218Z")

	testTable := []struct {
		name                string
		idCom               string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:  "OK",
			idCom: "7",
			mockBehavior: func(s *mock_comment.MockService, idCom int) {
				s.EXPECT().Get(idCom).Return(&comment.Comment{
					Id:     7,
					Author: 1,
					IdPost: 1,
					Time:   TimeTest,
					Body:   "Hey, I'm Alex",
				}, nil)
			},
			expectedStatusCode: 200,
			expectedRequestBody: `{"comment":{	"id":7,
												"author":1,
												"idPost":1,
												"time":"2023-08-03T07:27:50.35218Z",
												"body":"Hey, I'm Alex"
												}
											}`,
		},
		{
			name:  "foreign id",
			idCom: "0",
			mockBehavior: func(s *mock_comment.MockService, idCom int) {
				s.EXPECT().Get(idCom).Return(nil, errors.New("you can't change this comment"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"error": "you can't change this comment"}`,
		},
		{
			name:  "not int in url",
			idCom: "lol",
			mockBehavior: func(s *mock_comment.MockService, idCom int) {
			},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error": "you can't change this comment"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			idComment, _ := strconv.Atoi(testCase.idCom)
			comService := mock_comment.NewMockService(ctrl)
			testCase.mockBehavior(comService, idComment)

			Handler := &handler{Service: comService}

			r := gin.Default()
			r.GET("/blog/post/comment/:idCom", Handler.Get())

			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", "/blog/post/comment/"+testCase.idCom, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			require.JSONEq(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_Create(t *testing.T) {
	type mockBehavior func(s *mock_comment.MockService, com *comment.Comment)

	testTable := []struct {
		name                string
		idPost              string
		inputBody           string `json:"inputBody"`
		inputComment        *comment.Comment
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			idPost:    "1",
			inputBody: `{ "body": "Yop!"}`,
			inputComment: &comment.Comment{
				Id:     0,
				Author: 1,
				Time:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				IdPost: 1,
				Body:   "Yop!",
			},
			mockBehavior: func(s *mock_comment.MockService, com *comment.Comment) {
				s.EXPECT().Create(com).Return(10, nil)
			},
			expectedStatusCode:  http.StatusCreated,
			expectedRequestBody: `{"message": "comment is created"}`,
		},
		{
			name:      "not int in url",
			idPost:    "kek",
			inputBody: `{ "body": "Yop!"}`,
			inputComment: &comment.Comment{
				Id:     0,
				Author: 1,
				Time:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				IdPost: 1,
				Body:   "Yop!",
			},
			mockBehavior: func(s *mock_comment.MockService, com *comment.Comment) {
			},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error": "not int in url param"}`,
		},
		{
			name:      "bad request body",
			idPost:    "1",
			inputBody: `{ "bod": "Yop!"}`,
			inputComment: &comment.Comment{
				Id:     0,
				Author: 1,
				Time:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				IdPost: 1,
				Body:   "Yop!",
			},
			mockBehavior: func(s *mock_comment.MockService, com *comment.Comment) {
			},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error": "bad request body"}`,
		},
		{
			name:      "Internal Error",
			idPost:    "1",
			inputBody: `{ "body": "Yop!"}`,
			inputComment: &comment.Comment{
				Id:     0,
				Author: 1,
				Time:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				IdPost: 1,
				Body:   "Yop!",
			},
			mockBehavior: func(s *mock_comment.MockService, com *comment.Comment) {
				s.EXPECT().Create(com).Return(0, errors.New("example error"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"error": "internal server error"}`,
		},
	}

	defer func(orig func() time.Time) {
		timeNow = time.Now
	}(timeNow)

	timeNow = func() time.Time {
		return time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			comService := mock_comment.NewMockService(ctrl)
			testCase.mockBehavior(comService, testCase.inputComment)

			Handler := &handler{Service: comService}

			r := gin.Default()

			r.POST("/blog/post/:idPost/comment", func(c *gin.Context) {
				c.Set("userId", 1)
				c.Next()
			}, Handler.Create())

			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/blog/post/"+testCase.idPost+"/comment",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			//require.JSONEq(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}

}
