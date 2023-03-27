package dpfm_api_output_formatter

import (
	"data-platform-api-quantity-unit-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToQuantityUnit(rows *sql.Rows) (*[]QuantityUnit, error) {
	defer rows.Close()
	quantityUnit := make([]QuantityUnit, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuantityUnit{}

		err := rows.Scan(
			&pm.QuantityUnit,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &quantityUnit, nil
		}

		data := pm
		quantityUnit = append(quantityUnit, QuantityUnit{
			QuantityUnit: data.QuantityUnit,
		})
	}

	return &quantityUnit, nil
}

func ConvertToQuantityUnitText(rows *sql.Rows) (*[]QuantityUnitText, error) {
	defer rows.Close()
	quantityUnitText := make([]QuantityUnitText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuantityUnitText{}

		err := rows.Scan(
			&pm.QuantityUnit,
			&pm.Language,
			&pm.QuantityUnitName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &quantityUnitText, err
		}

		data := pm
		quantityUnitText = append(quantityUnitText, QuantityUnitText{
			QuantityUnit:     data.QuantityUnit,
			Language:         data.Language,
			QuantityUnitName: data.QuantityUnitName,
		})
	}

	return &quantityUnitText, nil
}

func ConvertToQuantityUnitTexts(rows *sql.Rows) (*[]QuantityUnitText, error) {
	defer rows.Close()
	quantityUnitText := make([]QuantityUnitText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuantityUnitTexts{}

		err := rows.Scan(
			&pm.QuantityUnit,
			&pm.Language,
			&pm.QuantityUnitName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &quantityUnitText, err
		}

		data := pm
		quantityUnitText = append(quantityUnitText, QuantityUnitText{
			QuantityUnit:     data.QuantityUnit,
			Language:         data.Language,
			QuantityUnitName: data.QuantityUnitName,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &quantityUnitText, nil
	}

	return &quantityUnitText, nil
}
