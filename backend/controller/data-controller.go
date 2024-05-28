package controller

import (
	"backend/entity"
	"backend/service"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/jszwec/csvutil"
)

type DataController interface {
	Start()
}

type dataController struct{}

func NewDataController() DataController {
	return &dataController{}
}

func (ctrl *dataController) Start() {
	insertType()
	insertEvent()
	insertProduct()
	insertNews()
	insertProductEvent()
}

func insertType() {
	fileName := "data/type-data.csv"
	file, dcr := getData(fileName)
	if file != nil {
		defer file.Close()
		if dcr != nil {
			var itms []entity.Type
			err := dcr.Decode(&itms)
			if err != nil {
				panic("Failed to parsing to Type object")
			}
			svc := service.NewTypeService()
			for _, itm := range itms {
				svc.Insert(&itm)
			}
		}
	}
}

func insertEvent() {
	fileName := "data/event-data.csv"
	file, dcr := getData(fileName)
	if file != nil {
		defer file.Close()
		fmt.Println(dcr)
		if dcr != nil {
			var itms []entity.Event
			err := dcr.Decode(&itms)
			if err != nil {
				panic("Failed to parsing to Event object")
			}
			svc := service.NewEventService()
			for _, itm := range itms {
				svc.Insert(&itm)
			}
		}
	}
}

func insertProduct() {
	fileName := "data/product-data.csv"
	file, dcr := getData(fileName)
	if file != nil {
		defer file.Close()
		if dcr != nil {
			var itms []entity.Product
			err := dcr.Decode(&itms)
			if err != nil {
				panic("Failed to parsing to Product object")
			}
			svc := service.NewProductService()
			for _, itm := range itms {
				svc.Insert(&itm)
			}
		}
	}
}

func insertNews() {
	fileName := "data/news-data.csv"
	file, dcr := getData(fileName)
	if file != nil {
		defer file.Close()
		if dcr != nil {
			var itms []entity.News
			err := dcr.Decode(&itms)
			if err != nil {
				panic("Failed to parsing to News object")
			}
			svc := service.NewNewsService()
			for _, itm := range itms {
				svc.Insert(&itm)
			}
		}
	}
}

func insertProductEvent() {
	fileName := "data/product-event-data.csv"
	file, err := os.Open(fileName)
	if err != nil {
		panic("Failed to opening file")
	}
	rdr := csv.NewReader(file)
	// using FieldsPerRecord to a negative for read csv file with different columns of rows
	rdr.FieldsPerRecord = -1
	itms, err := rdr.ReadAll()
	if err != nil {
		panic("Failed to read all file")
	}
	pdtSvc := service.NewProductService()
	for i := 1; i < len(itms); i++ {
		pdtId := parseID(itms[i][0])
		pdt := pdtSvc.Read(pdtId)
		evtLen := len(itms[i])
		evtLst := make([]entity.Event, evtLen-1)
		for j := 1; j < evtLen; j++ {
			evtId := parseID(itms[i][j])
			evtLst[j-1] = entity.Event{ID: evtId}
		}
		pdt.Events = &evtLst
		pdtSvc.Insert(pdt)
	}
}

func getData(fileName string) (*os.File, *csvutil.Decoder) {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Failed to opening file")
	}
	reader := csv.NewReader(file)
	dcr, err := csvutil.NewDecoder(reader)
	if err != nil {
		panic("Failed to decoding data")
	}
	return file, dcr
}

func parseID(id string) uint32 {
	res, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic("Failed to parsing id")
	}
	return uint32(res)
}
