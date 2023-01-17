package gtfs

import (
	"gtfs_viewer/src/helpers"
	"strings"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func movingStopsRoute(context *gin.Context) {
	area := context.Param("area")

    dateParam := context.Query("date")
	bounds := context.Query("bounds")

    if dateParam == "" {
		context.String(http.StatusBadRequest, "Param 'date' is missing")
		return
    } else {

		boundsStrings := strings.Split(bounds, ",")
		var boundsValues [4]float32
		for index, element := range boundsStrings {
			element, _ := strconv.ParseFloat(element, 64)
			boundsValues[index] = float32(element)
		}

		dataFound := SelectData(area)

		date, _ := strconv.Atoi(dateParam)
		// TODO add error condition check
		stopsFound := FilterByDate(dataFound.Data, uint32(date), boundsValues)

		context.JSON(http.StatusOK, stopsFound)
 	}	
}

func rangeDatesRoute(context *gin.Context) {
	area := context.Param("area")

	dataFound := SelectData(area)

	result := RangeDataModel{
		DataBounds: dataFound.Bounds,
		StartDate: dataFound.StartDate,
		EndDate: dataFound.EndDate,
	}
	context.JSON(http.StatusOK, result)
}

func transportTypeRoute(context *gin.Context) {
	area := context.Param("area")

	dataFound := SelectData(area)
	context.JSON(http.StatusOK, dataFound.routeTypes)
}

func availableAreasRoute(context *gin.Context) {
	var availableAreas []string
	for _, feature := range gtfsInputData.Files {
		availableAreas = append(availableAreas, feature.Title)
	}
	context.JSON(http.StatusOK, availableAreas)
}

func GtfsGroupRouterHandler(dataPath string, router *gin.Engine) {
	
	// get data
	gtfsSuffix := "_gtfsData.json"
	gtfsInputData = GetData(dataPath, gtfsSuffix)
	helpers.PrintMemresultUsage()

	group := router.Group("/api/v2/gtfs_builder")

	group.GET(":area/moving_nodes", movingStopsRoute)
	group.GET(":area/range_dates", rangeDatesRoute)
	group.GET(":area/route_types", transportTypeRoute)
	group.GET("/existing_study_areas", availableAreasRoute)
}