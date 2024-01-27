local yab = require("yab")
require("env")

-- bundle resources
yab.task(yab.find("./resources/**"), "./view/resources.go", function()
	os.execute("rm -f ./view/resources.go")
	os.execute("fyne bundle -package view -o ./view/resources.go ./resources")
end)
