package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	name := "name"
	reference := "reference"
	priceMin := float64(1)

	filter := ProductFilter{
		Name:            &name,
		Reference:       &reference,
		Status:          []string{"status"},
		Category:        nil,
		StockCity:       nil,
		Supplier:        nil,
		PriceMin:        &priceMin,
		PriceMax:        nil,
		AvailableMin:    nil,
		AvailableMax:    nil,
		DateFrom:        nil,
		DateTo:          nil,
		BeforeCreatedAt: nil,
		Limit:           0,
	}

	v := reflect.ValueOf(filter)
	t := reflect.TypeOf(filter)
	numField := t.NumField()

	for i := 0; i < numField; i++ {
		field := t.Field(i)
		value := v.Field(i)

		sqlCond := field.Tag.Get("query")
		if sqlCond == "" || sqlCond == "-" {
			continue
		}

		if value.Kind() == reflect.Ptr && !value.IsNil() {
			if strings.Contains(sqlCond, "ILIKE") {
				aaa := fmt.Sprintf("%s%%", value.Elem().Interface())
				fmt.Println(aaa)
			} else {
				fmt.Println(sqlCond, value.Elem().Interface())
			}
		}

		if value.Kind() == reflect.Slice && value.Len() > 0 {
			fmt.Println(sqlCond, value.Interface())
		}
	}
}

type ProductFilter struct {
	Name            *string    `json:"name" query:"name ILIKE ?"`
	Reference       *string    `json:"reference" query:"reference = ?"`
	Status          []string   `json:"status" query:"status in ?"`
	Category        []string   `json:"category" query:"category_id in ?"`
	StockCity       []string   `json:"stock_city" query:"stock_city in ?"`
	Supplier        *string    `json:"supplier" query:"supplier_id = ?"`
	PriceMin        *float64   `json:"price_min" query:"price >= ?"`
	PriceMax        *float64   `json:"price_max" query:"price <= ?"`
	AvailableMin    *int       `json:"available_min" query:"quantity >= ?"`
	AvailableMax    *int       `json:"available_max" query:"quantity <= ?"`
	DateFrom        *time.Time `json:"date_from" query:"added_date >= ?"`
	DateTo          *time.Time `json:"date_to" query:"added_date <= ?"`
	BeforeCreatedAt *time.Time `json:"before_created_at" query:"created_at > ?"`
	Limit           int        `json:"limit" query:"-"`
}
