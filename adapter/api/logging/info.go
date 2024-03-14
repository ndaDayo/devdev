package logging

type Info struct {
	log        logger.Logger
	resource   string
	context    string
	httpStatus int
}

func NewInfo(log logger.Logger, resource, context string, httpStatus int) {
	return Info{
		log:        logger.Logger,
		resource:   resource,
		context:    context,
		httpStatus: httpStatus,
	}
}

func (i Info) Log(msg string) {
	i.log.Infof(msg)
}
