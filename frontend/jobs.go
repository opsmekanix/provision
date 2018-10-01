package frontend

import (
	"fmt"
	"net/http"

	"github.com/VictorLowther/jsonpatch2"
	"github.com/digitalrebar/provision/backend"
	"github.com/digitalrebar/provision/models"
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
)

// NOTE: Jobs are restricted by Machine UUID

// JobResponse return on a successful GET, PUT, PATCH or POST of a single Job
// swagger:response
type JobResponse struct {
	// in: body
	Body *models.Job
}

// JobsResponse return on a successful GET of all Jobs
// swagger:response
type JobsResponse struct {
	// in: body
	Body []*models.Job
}

// JobActionsResponse return on a successful GET of a Job's actions
// swagger:response
type JobActionsResponse struct {
	// in: body
	Body models.JobActions
}

// JobParamsResponse return on a successful GET of all Job's Params
// swagger:response
type JobParamsResponse struct {
	// in: body
	Body map[string]interface{}
}

// This is a HACK - I can't figure out how to get
// swagger to render this a binary.  So we lie.
// We also override this object from the server
// directory to have a binary format which
// turns it into a stream.
//
// JobLogResponse returned on a successful GET of a log
// swagger:response
type JobLogResponse struct {
	// in: body
	// format: binary
	Body string
}

// JobBodyParameter used to inject a Job
// swagger:parameters createJob putJob
type JobBodyParameter struct {
	// in: body
	// required: true
	Body *models.Job
}

// JobPatchBodyParameter used to patch a Job
// swagger:parameters patchJob
type JobPatchBodyParameter struct {
	// in: body
	// required: true
	Body jsonpatch2.Patch
}

// JobLogBodyParameter used to append to a Job log
// swagger:parameters putJobLog
type JobLogPutBodyParameter struct {
	// in: body
	// required: true
	Body interface{}
}

// JobPathParameter used to find a Job in the path
// swagger:parameters putJobs getJob putJob patchJob deleteJob getJobParams postJobParams getJobActions getJobLog putJobLog headJob
type JobPathParameter struct {
	// in: path
	// required: true
	// swagger:strfmt uuid
	Uuid uuid.UUID `json:"uuid"`
}

// JobParamsBodyParameter used to set Job Params
// swagger:parameters postJobParams
type JobParamsBodyParameter struct {
	// in: body
	// required: true
	Body map[string]interface{}
}

// JobListPathParameter used to limit lists of Job by path options
// swagger:parameters listJobs listStatsJobs
type JobListPathParameter struct {
	// in: query
	Offest int `json:"offset"`
	// in: query
	Limit int `json:"limit"`
	// in: query
	Available string
	// in: query
	Valid string
	// in: query
	ReadOnly string
	// in: query
	Uuid string
	// in: query
	Stage string
	// in: query
	Task string
	// in: query
	State string
	// in: query
	Machine string
	// in: query
	Archived string
	// in: query
	StartTime string
	// in: query
	EndTime string
}

// JobActionsPathParameter used to find a Job / Actions in the path
// swagger:parameters getJobActions
type JobActionsPathParameter struct {
	// in: path
	// required: true
	// swagger:strfmt uuid
	Uuid uuid.UUID `json:"uuid"`
	// in: query
	Plugin string `json:"plugin"`
	// in: query
	OS string `json:"os"`
}

// JobActionPathParameter used to find a Job / Action in the path
// swagger:parameters getJobAction
type JobActionPathParameter struct {
	// in: path
	// required: true
	// swagger:strfmt uuid
	Uuid uuid.UUID `json:"uuid"`
	// in: path
	// required: true
	Cmd string `json:"cmd"`
	// in: query
	Plugin string `json:"plugin"`
}

// JobActionBodyParameter used to post a Job / Action in the path
// swagger:parameters postJobAction
type JobActionBodyParameter struct {
	// in: path
	// required: true
	// swagger:strfmt uuid
	Uuid uuid.UUID `json:"uuid"`
	// in: path
	// required: true
	Cmd string `json:"cmd"`
	// in: query
	Plugin string `json:"plugin"`
	// in: body
	// required: true
	Body map[string]interface{}
}

func (f *Frontend) InitJobApi() {
	// swagger:route GET /jobs Jobs listJobs
	//
	// Lists Jobs filtered by some parameters.
	//
	// This will show all Jobs by default.
	//
	// You may specify:
	//    Offset = integer, 0-based inclusive starting point in filter data.
	//    Limit = integer, number of items to return
	//
	// Functional Indexs:
	//    Uuid = string
	//    Stage = string
	//    Task = string
	//    State = string
	//    Machine = string
	//    Archived = boolean
	//    StartTime = datetime
	//    EndTime = datetime
	//    Available = boolean
	//    Valid = boolean
	//    ReadOnly = boolean
	//
	// Functions:
	//    Eq(value) = Return items that are equal to value
	//    Lt(value) = Return items that are less than value
	//    Lte(value) = Return items that less than or equal to value
	//    Gt(value) = Return items that are greater than value
	//    Gte(value) = Return items that greater than or equal to value
	//    Between(lower,upper) = Return items that are inclusively between lower and upper
	//    Except(lower,upper) = Return items that are not inclusively between lower and upper
	//
	// Example:
	//    Uuid=fred - returns items named fred
	//    Uuid=Lt(fred) - returns items that alphabetically less than fred.
	//    Uuid=Lt(fred)&Archived=true - returns items with Uuid less than fred and Archived is true
	//
	// Responses:
	//    200: JobsResponse
	//    401: NoContentResponse
	//    403: NoContentResponse
	//    406: ErrorResponse
	f.ApiGroup.GET("/jobs",
		func(c *gin.Context) {
			f.List(c, &backend.Job{})
		})

	// swagger:route HEAD /jobs Jobs listStatsJobs
	//
	// Stats of the List Jobs filtered by some parameters.
	//
	// This will return headers with the stats of the list.
	//
	// You may specify:
	//    Offset = integer, 0-based inclusive starting point in filter data.
	//    Limit = integer, number of items to return
	//
	// Functional Indexs:
	//    Uuid = string
	//    Stage = string
	//    Task = string
	//    State = string
	//    Machine = string
	//    Archived = boolean
	//    StartTime = datetime
	//    EndTime = datetime
	//    Available = boolean
	//    Valid = boolean
	//    ReadOnly = boolean
	//
	// Functions:
	//    Eq(value) = Return items that are equal to value
	//    Lt(value) = Return items that are less than value
	//    Lte(value) = Return items that less than or equal to value
	//    Gt(value) = Return items that are greater than value
	//    Gte(value) = Return items that greater than or equal to value
	//    Between(lower,upper) = Return items that are inclusively between lower and upper
	//    Except(lower,upper) = Return items that are not inclusively between lower and upper
	//
	// Example:
	//    Uuid=fred - returns items named fred
	//    Uuid=Lt(fred) - returns items that alphabetically less than fred.
	//    Uuid=Lt(fred)&Archived=true - returns items with Uuid less than fred and Archived is true
	//
	// Responses:
	//    200: NoContentResponse
	//    401: NoContentResponse
	//    403: NoContentResponse
	//    406: ErrorResponse
	f.ApiGroup.HEAD("/jobs",
		func(c *gin.Context) {
			f.ListStats(c, &backend.Job{})
		})

	// swagger:route POST /jobs Jobs createJob
	//
	// Create a Job
	//
	// Create a Job from the provided object, Only Machine and UUID are used.
	//
	//     Responses:
	//       201: JobResponse
	//       202: JobResponse
	//       204: NoContentResponse
	//       400: ErrorResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       409: ErrorResponse
	//       422: ErrorResponse
	//       500: ErrorResponse
	f.ApiGroup.POST("/jobs",
		func(c *gin.Context) {
			// We don't use f.Create() because we need to be able to assign random
			// UUIDs to new Jobs without forcing the client to do so, yet allow them
			// for testing purposes amd if they alrady have a UUID scheme for jobs.
			b := &backend.Job{}
			if !assureDecode(c, b) {
				return
			}
			if b.Machine == nil {
				c.JSON(http.StatusBadRequest, models.NewError(c.Request.Method, http.StatusBadRequest, "Create request must have Machine field"))
				return
			}
			if !f.assureSimpleAuth(c, "jobs", "create", b.AuthKey()) {
				return
			}
			if b.Uuid == nil || len(b.Uuid) == 0 {
				b.Uuid = uuid.NewRandom()
			}
			err := &models.Error{
				Code: http.StatusNoContent,
			}
			var thunk func(*models.Error) *backend.Job
			rt := f.rt(c, b.Locks("create")...)
			rt.Do(func(_ backend.Stores) {
				b, thunk = realCreateJob(f, rt, b, err)
			})
			if thunk != nil {
				b = thunk(err)
			}
			switch err.Code {
			case http.StatusAccepted, http.StatusCreated:
				s, ok := models.Model(b.Job).(Sanitizable)
				var res *models.Job
				if ok {
					res = s.Sanitize().(*models.Job)
				} else {
					res = b.Job
				}
				c.JSON(err.Code, res)
			case http.StatusNoContent:
				c.Data(err.Code, gin.MIMEJSON, nil)
			default:
				c.JSON(err.Code, err)
			}
		})

	// swagger:route GET /jobs/{uuid} Jobs getJob
	//
	// Get a Job
	//
	// Get the Job specified by {uuid} or return NotFound.
	//
	//     Responses:
	//       200: JobResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	f.ApiGroup.GET("/jobs/:uuid",
		func(c *gin.Context) {
			f.Fetch(c, &backend.Job{}, c.Param(`uuid`))
		})

	// swagger:route HEAD /jobs/{uuid} Jobs headJob
	//
	// See if a Job exists
	//
	// Return 200 if the Job specifiec by {uuid} exists, or return NotFound.
	//
	//     Responses:
	//       200: NoContentResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: NoContentResponse
	f.ApiGroup.HEAD("/jobs/:uuid",
		func(c *gin.Context) {
			f.Exists(c, &backend.Job{}, c.Param(`uuid`))
		})

	// swagger:route PATCH /jobs/{uuid} Jobs patchJob
	//
	// Patch a Job
	//
	// Update a Job specified by {uuid} using a RFC6902 Patch structure
	//
	//     Responses:
	//       200: JobResponse
	//       400: ErrorResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       406: ErrorResponse
	//       422: ErrorResponse
	f.ApiGroup.PATCH("/jobs/:uuid",
		func(c *gin.Context) {
			f.Patch(c, &backend.Job{}, c.Param(`uuid`))
		})

	// swagger:route PUT /jobs/{uuid} Jobs putJob
	//
	// Put a Job
	//
	// Update a Job specified by {uuid} using a JSON Job
	//
	//     Responses:
	//       200: JobResponse
	//       400: ErrorResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       422: ErrorResponse
	f.ApiGroup.PUT("/jobs/:uuid",
		func(c *gin.Context) {
			f.Update(c, &backend.Job{}, c.Param(`uuid`))
		})

	// swagger:route DELETE /jobs/{uuid} Jobs deleteJob
	//
	// Delete a Job
	//
	// Delete a Job specified by {uuid}
	//
	//     Responses:
	//       200: JobResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       422: ErrorResponse
	f.ApiGroup.DELETE("/jobs/:uuid",
		func(c *gin.Context) {
			f.Remove(c, &backend.Job{}, c.Param(`uuid`))
		})

	// swagger:route GET /jobs/{uuid}/actions Jobs getJobActions
	//
	// Get actions for this job
	//
	// Get actions for the Job specified by {uuid} or return NotFound.
	//
	//     Responses:
	//       200: JobActionsResponse
	//       400: ErrorResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       409: ErrorResponse
	//       422: ErrorResponse
	f.ApiGroup.GET("/jobs/:uuid/actions",
		func(c *gin.Context) {
			uuid := c.Param(`uuid`)
			j := &backend.Job{}
			var bad bool
			var err error
			rt := f.rt(c, j.Locks("actions")...)
			rt.Do(func(d backend.Stores) {
				var jo models.Model
				if jo = rt.Find("jobs", uuid); jo == nil {
					err = &models.Error{Code: http.StatusNotFound, Type: backend.ValidationError,
						Messages: []string{fmt.Sprintf("Job %s does not exist", uuid)}}
					bad = true
					return
				}
				j = backend.AsJob(jo)
			})
			if bad {
				c.JSON(err.(*models.Error).Code, err)
				return
			}

			if !f.assureSimpleAuth(c, "jobs", "actions", j.AuthKey()) {
				return
			}
			rt.Debugf("Rendering jobs for %s os %s", uuid, c.Query("os"))
			actions, err := j.RenderActions(rt, c.Query("os"))
			if err != nil {
				be, ok := err.(*models.Error)
				if ok {
					c.JSON(be.Code, be)
				} else {
					c.JSON(http.StatusBadRequest, models.NewError(c.Request.Method, http.StatusBadRequest, err.Error()))
				}
				return
			}
			c.JSON(http.StatusOK, actions)

		})

	// swagger:route GET /jobs/{uuid}/log Jobs getJobLog
	//
	// Get the log for this job
	//
	// Get log for the Job specified by {uuid} or return NotFound.
	//
	//     Produces:
	//       application/octet-stream
	//       application/json
	//
	//     Responses:
	//       200: JobLogResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       404: ErrorResponse
	//       500: ErrorResponse
	f.ApiGroup.GET("/jobs/:uuid/log",
		func(c *gin.Context) {
			uuid := c.Param(`uuid`)
			j := &backend.Job{}
			var bad bool
			var err *models.Error
			var path string
			rt := f.rt(c, j.Locks("get")...)
			rt.Do(func(d backend.Stores) {
				var jo models.Model
				if jo = rt.Find("jobs", uuid); jo == nil {
					err = &models.Error{Code: http.StatusNotFound, Type: backend.ValidationError,
						Messages: []string{fmt.Sprintf("Job %s does not exist", uuid)}}
					bad = true
					return
				}
				j = backend.AsJob(jo)
				path = j.LogPath(rt)
			})
			if bad {
				c.JSON(err.Code, err)
				return
			}

			if !f.assureSimpleAuth(c, "jobs", "log", j.AuthKey()) {
				return
			}

			c.Writer.Header().Set("Content-Type", "application/octet-stream")
			c.File(path)
		})

	// swagger:route PUT /jobs/{uuid}/log Jobs putJobLog
	//
	// Append the string to the end of the job's log.
	//     Consumes:
	//       application/octet-stream
	//
	//     Produces:
	//       application/json
	//
	//     Responses:
	//       204: NoContentResponse
	//       400: ErrorResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	//       415: ErrorResponse
	//       404: ErrorResponse
	//       500: ErrorResponse
	f.ApiGroup.PUT("/jobs/:uuid/log",
		func(c *gin.Context) {
			if c.Request.Body == nil {
				err := &models.Error{Code: http.StatusBadRequest}
				c.JSON(err.Code, err)
				return
			}
			defer c.Request.Body.Close()
			if c.Request.Header.Get(`Content-Type`) != `application/octet-stream` {
				c.JSON(http.StatusUnsupportedMediaType,
					models.NewError("API ERROR", http.StatusUnsupportedMediaType,
						"job log put must have content-type application/octet-stream"))
				return
			}
			uuid := c.Param(`uuid`)
			j := &backend.Job{}
			var bad bool
			var err *models.Error
			rt := f.rt(c, j.Locks("get")...)
			rt.Do(func(d backend.Stores) {
				var jo models.Model
				if jo = d("jobs").Find(uuid); jo == nil {
					err = &models.Error{Code: http.StatusNotFound, Type: backend.ValidationError,
						Messages: []string{fmt.Sprintf("Job %s does not exist", uuid)}}
					bad = true
					return
				}
				j = backend.AsJob(jo)
			})
			if bad {
				c.JSON(err.Code, err)
				return
			}

			if !f.assureSimpleAuth(c, "jobs", "log", j.AuthKey()) {
				return
			}

			rt.Do(func(d backend.Stores) {
				var jo models.Model
				if jo = d("jobs").Find(uuid); jo == nil {
					err = &models.Error{Code: http.StatusNotFound, Type: backend.ValidationError,
						Messages: []string{fmt.Sprintf("Job %s does not exist", uuid)}}
					bad = true
					return
				}
				j = backend.AsJob(jo)

				if err := j.Log(rt, c.Request.Body); err != nil {
					err2 := &models.Error{Code: http.StatusInternalServerError, Type: "Server ERROR",
						Messages: []string{err.Error()}}
					c.JSON(err2.Code, err2)
				} else {
					c.Data(http.StatusNoContent, gin.MIMEJSON, nil)
				}
			})
			if bad {
				c.JSON(err.Code, err)
				return
			}
		})

	job := &backend.Job{}
	pActions, pAction, pRun := f.makeActionEndpoints(job.Prefix(), job, "uuid")

	// swagger:route GET /jobs/{uuid}/plugin_actions Jobs getJobActions
	//
	// List job plugin_actions Job
	//
	// List Job plugin_actions for a Job specified by {uuid}
	//
	// Optionally, a query parameter can be used to limit the scope to a specific plugin.
	//   e.g. ?plugin=fred
	//
	//     Responses:
	//       200: ActionsResponse
	//       401: NoJobResponse
	//       403: NoJobResponse
	//       404: ErrorResponse
	f.ApiGroup.GET("/jobs/:uuid/plugin_actions", pActions)

	// swagger:route GET /jobs/{uuid}/plugin_actions/{cmd} Jobs getJobAction
	//
	// List specific action for a job Job
	//
	// List specific {cmd} action for a Job specified by {uuid}
	//
	// Optionally, a query parameter can be used to limit the scope to a specific plugin.
	//   e.g. ?plugin=fred
	//
	//     Responses:
	//       200: ActionResponse
	//       400: ErrorResponse
	//       401: NoJobResponse
	//       403: NoJobResponse
	//       404: ErrorResponse
	f.ApiGroup.GET("/jobs/:uuid/plugin_actions/:cmd", pAction)

	// swagger:route POST /jobs/{uuid}/plugin_actions/{cmd} Jobs postJobAction
	//
	// Call an action on the node.
	//
	// Optionally, a query parameter can be used to limit the scope to a specific plugin.
	//   e.g. ?plugin=fred
	//
	//     Responses:
	//       400: ErrorResponse
	//       200: ActionPostResponse
	//       401: NoJobResponse
	//       403: NoJobResponse
	//       404: ErrorResponse
	//       409: ErrorResponse
	f.ApiGroup.POST("/jobs/:uuid/plugin_actions/:cmd", pRun)
}
