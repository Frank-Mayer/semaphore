require("env")
require("resources")

local yab = require("yab")

local ost = yab.os_type()
if ost == "windows" then
	yab.task(yab.find("**.go"), "Semaphore.exe", function()
		os.execute('go build -ldflags="-s -w -H=windowsgui" -o Semaphore.exe')
	end)
	os.execute("Semaphore.exe")
else -- unix
	yab.task(yab.find("**.go"), "Semaphore.a", function()
		os.execute('go build -ldflags="-s -w" -o Semaphore')
	end)
	os.execute("./Semaphore")
end
