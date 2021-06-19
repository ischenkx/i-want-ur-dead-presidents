package util

type ProductValidator struct {}

//func (v *ProductValidator) ValidateUpdate(productDto entities.UpdateEntityDto) error {
//	if productDto.OwnerID != nil {
//		if _, err := uuid.Parse(*productDto.OwnerID); err != nil {
//			return err
//		}
//	}
//
//	if productDto.MoneyGoal != nil {
//		if *productDto.MoneyGoal < 0 {
//			return errors.New("money goal must be not negative")
//		}
//	}
//
//	return nil
//}
//
//func (v *ProductValidator) ValidateCreate(product models.Product) error {
//	if _, err := uuid.Parse(product.OwnerID); err != nil {
//		return err
//	}
//
//	if product.MoneyGoal < 0 {
//			return errors.New("money goal must be not negative")
//		}
//
//	return nil
//}
