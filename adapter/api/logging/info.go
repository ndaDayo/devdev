package logging

type Info struct {
	log      logger.Logger
	resource string
	context  string
}

func NewInfo(log logger.Logger, resource string, context string) Info {
	return Info{
		log:      log,
		resource: resource,
		context:  context,
	}
}
