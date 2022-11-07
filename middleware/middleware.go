package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/ArdyJunata/api-test-cakra/constants"
	"github.com/ArdyJunata/api-test-cakra/pkg/logger"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
	log logger.Logger
}

func NewMiddleware(log logger.Logger) *Middleware {
	return &Middleware{
		log: log,
	}
}

func (m *Middleware) Tracer(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var data = map[logger.LogKey]interface{}{}

		now := time.Now()
		tracer := r.Header.Get(constants.TRACER_ID)
		if tracer == "" {
			tracer = uuid.New().String()
		}
		ctx := r.Context()

		data[logger.TRACER_ID] = tracer

		ctx = context.WithValue(ctx, logger.TRACER_ID, tracer)
		r = r.WithContext(ctx)

		method := r.Method
		path := r.URL.Path

		next(w, r, p)

		end := time.Since(now).Seconds()

		data[logger.RESPONSE_TIME] = end
		data[logger.METHOD] = method
		data[logger.RESPONSE_TYPE] = "seconds"
		data[logger.PATH] = path

		w.Header().Set("X-Trace-ID", tracer)

		ctx = context.WithValue(ctx, logger.DATA, data)

		m.log.Infof(ctx, "finished request")
	}
}
