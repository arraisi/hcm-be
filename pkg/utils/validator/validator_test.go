package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type MockStruct struct {
	Id int64 `json:"id" validate:"required"`
}

func TestValidator(t *testing.T) {
	t.Run("ValidateStruct", func(t *testing.T) {
		req := MockStruct{
			Id: 1,
		}
		gotErr := ValidateStruct(req)
		require.NoError(t, gotErr)
	})

	t.Run("ValidateStruct error", func(t *testing.T) {
		req := MockStruct{}
		gotErr := ValidateStruct(req)
		require.Error(t, gotErr)
	})

	t.Run("ValidatePointerValue", func(t *testing.T) {
		valStr := "string"
		valInt := int64(1)
		valInt32 := int32(1)
		valSliceint64 := []int64{valInt}
		valSliceStr := []string{valStr}

		okStr := IsNotNilAndNotEmptyString(&valStr)
		okStr2 := IsNotEmptyString(valStr)
		okInt := IsNotNilAndNotZeroInt64(&valInt)
		okInt32 := IsNotNilAndNotZeroInt32(&valInt32)
		okSliceInt64 := IsNotEmptySlice(valSliceint64)
		okSliceStr := IsNotEmptySlice(valSliceStr)

		require.True(t, okStr)
		require.True(t, okStr2)
		require.True(t, okInt)
		require.True(t, okInt32)
		require.True(t, okSliceInt64)
		require.True(t, okSliceStr)
	})
}
