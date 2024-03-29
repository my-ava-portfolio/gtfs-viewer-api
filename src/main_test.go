package main

import (
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	gtfs "gtfs_viewer/src/core/stops"
	gtfsRoutes "gtfs_viewer/src/routers/gtfs"
)

func TestMovingNodesRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/fake/moving_nodes?date=1167642440&bounds=-180.0,-89.0,180.0,89.0", nil)
	Router.ServeHTTP(w, req)

    var stops []gtfs.StopItemFiltered
    json.Unmarshal(w.Body.Bytes(), &stops)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, stops)
}

func TestRangeDatesRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/fake/range_dates", nil)
	Router.ServeHTTP(w, req)

    var rangeData gtfs.RangeDataModel
    json.Unmarshal(w.Body.Bytes(), &rangeData)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, rangeData)
}

func TestRouteTypesRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/fake/route_types", nil)
	Router.ServeHTTP(w, req)

    var routeTypes []uint8
    json.Unmarshal(w.Body.Bytes(), &routeTypes)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, routeTypes)
}

func TestRouteLongNameRouteValid(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/fake/route_long_name?id=0", nil)
	Router.ServeHTTP(w, req)

    var RouteLongName string
    json.Unmarshal(w.Body.Bytes(), &RouteLongName)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, "Stagecoach - Airport Shuttle", RouteLongName)
}

func TestRouteLongNameRouteNotValid(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/fake/route_long_name?id=100", nil)
	Router.ServeHTTP(w, req)

    var RouteLongName string
    json.Unmarshal(w.Body.Bytes(), &RouteLongName)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, "missing", RouteLongName)
}

func TestAvailableAreasRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/existing_study_areas", nil)
	Router.ServeHTTP(w, req)

    var availableAreas []string
    json.Unmarshal(w.Body.Bytes(), &availableAreas)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, availableAreas)
}

var Router *gin.Engine
func TestMain(m *testing.M)() {

	Router = setupRouter()
	gtfsRoutes.GtfsGroupRouterHandler("testData/", Router)

	// BEFORE tests
    exitVal := m.Run()
    // AFTER tests

    os.Exit(exitVal)
}