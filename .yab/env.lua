local yab = require("yab")

yab.use("golang", "1.21.6")

-- install fyne
os.execute("go install fyne.io/fyne/v2/cmd/fyne@latest")
