function decodeOnce(codeReader, selectedDeviceId) {
  codeReader
    .decodeFromInputVideoDevice(selectedDeviceId, "video")
    .then((result) => {
      console.log(result);
      // document.getElementById("result").textContent = result.text;
      console.log('result: ', result.text)
      var raw = "";
      var length = 0;
      var format = 0;
      for (var i in result.rawBytes) {
        i = parseInt(i);
        // check if we are not in the last byte
        if (i == result.rawBytes.length - 1) {
          break;
        }
        var val0 = toHex(result.rawBytes[i]);
        var val1 = toHex(result.rawBytes[i + 1]);
        // Get QR Code Format
        if (i == 0) {
          format = val0[0];
          length = "0x" + val0[1];
          continue;
        } else if (i == 1) {
          length += [val0[0]];
          length = parseInt(length);
        }
        // console.log(i, val0, val1);
        var tempByte = "0x" + val0[1] + val1[0];
        if (i != result.text.length - 1) {
          tempByte += ", ";
        }
        var c = parseInt(i);
        // console.log(i,c,c%10);
        if (c % 10 == 0 && i != 0) {
          tempByte += "\n";
        }
        raw += tempByte;
      }
      // document.getElementById("raw_code").textContent = raw.split("0x").join("");
      // $("#raw_code").text(raw.split("0x").join(""))  
      console.log("raw: ", raw.toString());
      console.log("formatted: ", raw.split("0x").join(""));
      console.log("raw: ", raw.split("0x").join(""))

      var pathname = window.location.pathname
      console.log("pathname:",pathname)

      $.ajax({
        method:"POST",
        url: "https://decoder.pudding.ws/api/decode",
        data: {
          "encoded_string": raw.split("0x").join(""),
        },success: function(response){
          console.log("response:",response)
          var v = JSON.parse(response)

          if (v.Status == "Error"){
            console.error("error:", v.Error)
          }

          if (v.Status == "Success"){
            var dataAction = localStorage.getItem("data-action")
            
            //* Sign up
            if (pathname == "/" && dataAction == "signup"){
              $("#reg-username").val(v.UID)
              $(".modal-container").removeClass("open")

              $("#register-btn").on("click", function(){
                if ($("#reg-username").val() != "" && $("#reg-password").val() != "" && $("#reg-firstname").val() != "" && $("#reg-lastname").val() != "" && $("#reg-contact").val() != ""){
                  $.ajax({
                    method: "POST",
                    url: "/api/create_user",
                    data: {
                      "uid": v.UID,
                      "username": v.UID,
                      "password": $("#reg-password").val(),
                      "first_name": $("#reg-firstname").val(),
                      "last_name": $("#reg-lastname").val(),
                      "contact": $("#reg-contact").val()
                    },success: function(){
                      localStorage.removeItem("data-action")
                      window.location.reload()
                    }
                  })
                }
              })
            }

            //* Log in
            if (pathname == "/" && dataAction == "scanqr-login"){
              $.ajax({
                method: "POST",
                url: "/api/scan_login",
                data: {
                  "username": v.UID,
                },success: function(){
                  localStorage.removeItem("data-action")
                  window.location.reload()
                }
              })
            }
          }
        }
      })
    })
    .catch((err) => {
      console.error(err);
      // document.getElementById("result").textContent = err;
    });
}

function toHex(v) {
  var val = String(v.toString(16));
  if (val.length == 1) {
    val = "0" + val;
  }
  return val;
}

function decodeContinuously(codeReader, selectedDeviceId) {
  codeReader.decodeFromInputVideoDeviceContinuously(
    selectedDeviceId,
    "video",
    (result, err) => {
      if (result) {
        // properly decoded qr code
        console.log("Found QR code!", result);
        document.getElementById("result").textContent = result.text;
      }

      if (err) {
        // As long as this error belongs into one of the following categories
        // the code reader is going to continue as excepted. Any other error
        // will stop the decoding loop.
        //
        // Excepted Exceptions:
        //
        //  - NotFoundException
        //  - ChecksumException
        //  - FormatException

        if (err instanceof ZXing.NotFoundException) {
          console.log("No QR code found.");
        }

        if (err instanceof ZXing.ChecksumException) {
          console.log("A code was found, but it's read value was not valid.");
        }

        if (err instanceof ZXing.FormatException) {
          console.log("A code was found, but it was in a invalid format.");
        }
      }
    }
  );
}

window.addEventListener("load", function () {
  let selectedDeviceId;
  const codeReader = new ZXing.BrowserQRCodeReader();
  console.log("ZXing code reader initialized");

  codeReader
    .getVideoInputDevices()
    .then((videoInputDevices) => {
      const sourceSelect = document.getElementById("sourceSelect");
      selectedDeviceId = videoInputDevices[0].deviceId;
      if (videoInputDevices.length >= 1) {
        videoInputDevices.forEach((element) => {
          const sourceOption = document.createElement("option");
          sourceOption.text = element.label;
          sourceOption.value = element.deviceId;
          sourceSelect.appendChild(sourceOption);
        });

        sourceSelect.onchange = () => {
          selectedDeviceId = sourceSelect.value;
          decodeOnce(codeReader, selectedDeviceId);
        };

        const sourceSelectPanel = document.getElementById("sourceSelectPanel");
        sourceSelectPanel.style.display = "block";
      }

      document.getElementById("startButton").addEventListener("click", () => {
        const decodingStyle = document.getElementById("decoding-style").value;

        if (decodingStyle == "once") {
          decodeOnce(codeReader, selectedDeviceId);
        } else {
          decodeContinuously(codeReader, selectedDeviceId);
        }

        console.log(`Started decode from camera with id ${selectedDeviceId}`);
      });

      document.getElementById("resetButton").addEventListener("click", () => {
        codeReader.reset();
        document.getElementById("result").textContent = "";
        console.log("Reset.");
      });
    })
    .catch((err) => {
      console.error(err);
    });
});
