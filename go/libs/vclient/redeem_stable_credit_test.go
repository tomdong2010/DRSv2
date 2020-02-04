package vclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateRedeemStableCredit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "100", AssetCode: "vUSD"}
		err := redeemStableCreditInput.Validate()

		assert.Nil(t, err)
	})

	t.Run("fail, should throw error invalid RedeemAmount format", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "ABC", AssetCode: "vUSD"}
		err := redeemStableCreditInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "invalid RedeemAmount format", err.Error())
	})

	t.Run("fail, should throw error RedeemAmount must be positive", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "-2", AssetCode: "vUSD"}
		err := redeemStableCreditInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "RedeemAmount must be positive", err.Error())
	})

	t.Run("fail, should throw error assetCode must not be blank", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "100", AssetCode: ""}
		err := redeemStableCreditInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "assetCode must not be blank", err.Error())
	})

	t.Run("fail, should throw error invalid assetCode format", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{RedeemAmount: "100", AssetCode: "BAD_ASSET_CODE"}
		err := redeemStableCreditInput.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "invalid assetCode format", err.Error())
	})
}

func TestValidateRedeemStableCreditToAbiInput(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		redeemStableCreditInput := &RedeemStableCreditInput{AssetCode: "vUSD"}
		abiInput := redeemStableCreditInput.ToAbiInput()

		assert.NotNil(t, abiInput)
	})
}

//func TestRedeemStableCredit(t *testing.T) {
//	t.Run("success", func(t *testing.T) {
//		testHelper := testHelperWithMock(t)
//		defer testHelper.MockController.Finish()
//
//		pricePerAssetUnit, _ := utils.StringToAmount("1")
//
//		input := &RedeemStableCreditInput{
//			AssetCode: "vUSD",
//		}
//		abiInput := input.ToAbiInput()
//
//		testHelper.MockDRSContract.EXPECT().
//			GetExchange(gomock.AssignableToTypeOf(&bind.CallOpts{}), abiInput.AssetCode).
//			Return("vUSD", utils.StringToByte32("VELO"), pricePerAssetUnit, nil)
//
//		result, err := testHelper.Client.RedeemStableCredit(input)
//
//		assert.NoError(t, err)
//		assert.NotNil(t, result)
//		assert.Equal(t, "vUSD", result.AssetCode)
//		assert.Equal(t, utils.StringToByte32("VELO"), result.CollateralAssetCode)
//		assert.Equal(t, "10000000", result.PriceInCollateralPerAssetUnit.String())
//	})
//
//	t.Run("fail, should throw error assetCode must not be blank", func(t *testing.T) {
//		testHelper := testHelperWithMock(t)
//		defer testHelper.MockController.Finish()
//
//		input := &RedeemStableCreditInput{
//			AssetCode: "",
//		}
//
//		result, err := testHelper.Client.RedeemStableCredit(input)
//
//		assert.NotNil(t, err)
//		assert.Nil(t, result)
//		assert.Equal(t, "assetCode must not be blank", err.Error())
//	})
//}
