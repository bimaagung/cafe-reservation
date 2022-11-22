package usecase

import (
	"reflect"
	"testing"

	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/gofiber/fiber/v2"
)

func Test_menuUseCaseImpl_GetList(t *testing.T) {
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name         string
		useCase      *menuUseCaseImpl
		args         args
		wantResponse []domain.MenuRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResponse := tt.useCase.GetList(tt.args.ctx); !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("menuUseCaseImpl.GetList() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
