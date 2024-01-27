require("env")
require("resources")
local yab = require("yab")

local out
local ost = yab.os_type()
if ost == "windows" then
	out = "Semaphore.exe"
elseif ost == "darwin" then
	out = "Semaphore.app"
end

yab.task(yab.find("**.go"), out, function()
	os.execute("fyne package")
end)
