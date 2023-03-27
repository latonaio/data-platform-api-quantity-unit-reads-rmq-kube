package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-quantity-unit-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-quantity-unit-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var quantityUnit *[]dpfm_api_output_formatter.QuantityUnit
	var quantityUnitText *[]dpfm_api_output_formatter.QuantityUnitText
	for _, fn := range accepter {
		switch fn {
		case "QuantityUnit":
			func() {
				quantityUnit = c.QuantityUnit(mtx, input, output, errs, log)
			}()
		case "QuantityUnitText":
			func() {
				quantityUnitText = c.QuantityUnitText(mtx, input, output, errs, log)
			}()
		case "QuantityUnitTexts":
			func() {
				quantityUnitText = c.QuantityUnitTexts(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		QuantityUnit:     quantityUnit,
		QuantityUnitText: quantityUnitText,
	}

	return data
}

func (c *DPFMAPICaller) QuantityUnit(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.QuantityUnit {
	quantityUnit := input.QuantityUnit.QuantityUnit

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quantity_unit_quantity_unit_data
		WHERE QuantityUnit = ?;`, quantityUnit,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToQuantityUnit(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) QuantityUnitText(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.QuantityUnitText {
	var args []interface{}
	quantityUnit := input.QuantityUnit.QuantityUnit
	quantityUnitText := input.QuantityUnit.QuantityUnitText

	cnt := 0
	for _, v := range quantityUnitText {
		args = append(args, quantityUnit, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quantity_unit_text_data
		WHERE (QuantityUnit, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToQuantityUnitText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) QuantityUnitTexts(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.QuantityUnitText {
	var args []interface{}
	quantityUnitText := input.QuantityUnit.QuantityUnitText

	cnt := 0
	for _, v := range quantityUnitText {
		args = append(args, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?),", cnt-1) + "(?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quantity_unit_text_data
		WHERE Language IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	//
	data, err := dpfm_api_output_formatter.ConvertToQuantityUnitTexts(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
