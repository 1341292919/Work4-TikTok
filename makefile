MODULE = TikTok

SERVICE_NAME = video

.PHONY:target
target:
		sh build.sh
		sh output/bootstrap.sh

.PHONY:new
new:
	hz new \
	-module $(MODULE) \
	-service "$(SERVICE_NAME)"
	hz update -idl ./idl/interact.thrift
	hz update -idl ./idl/socialize.thrift
	hz update -idl ./idl/user.thrift
	hz update -idl ./idl/video.thrift
	hz update -idl ./idl/model.thrift

.PHONY: gen
gen:
	hz update -idl ./idl/interact.thrift
	hz update -idl ./idl/socialize.thrift
	hz update -idl ./idl/user.thrift
	hz update -idl ./idl/video.thrift
	hz update -idl ./idl/model.thrift