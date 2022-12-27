var tagBtn = document.querySelector(".add-tag__btn");
var dashboardModal = document.querySelector(".tag-add__modal");
var modalContainer = document.querySelector(".tag-modal__container");
var cancelBtn = document.querySelector("#tag-cancel");

function  createTag(){
  dashboardModal.classList.add('show');
  modalContainer.classList.add('show');
  $("#tag-save").attr("data-action", "save-tag")
  $("#timezone").empty()

  $(".employee-input").val("");
  $(".tag-header").html("Create Tag")
  $(".tag-header__desc").html("Fill out the form below to add.")
  $(".editing-img").css("display", "none")
  $("#machineID").css("pointer-events", "all").css("opacity", "1")

  $.get("/admin/api/d/timezone/read", function(index){
    var result = JSON.parse(index).result
    $("#timezone").append(`
      <option value="">Select time zone</option>
    `)

    $.each(result, function(_,tz){
      $("#timezone").append(`
        <option value="`+ tz.ID +`">`+ tz.Name +`</option>
      `)
    })
  })

  $("[data-action='save-tag']").on("click", function () {
    if ($("#machineID").val() != "" && $("#name").val() != "" && $("#description").val() != "" &&
      $("#address").val() != "" && $("longitude").val() != "" &&
      $("#latitude").val() != "" && $("#timezone").val() != ""){

      $.ajax({
        url: "/api/add_tag",
        data: {
          "uid": $("#machineID").val(),
          "code": $("#code").val(),
          "name": $("#name").val(),
          "description": $("#description").val(),
          "type": $("#maintenance-type").val(),
          "address": $("#address").val(),
          "longitude": $("#longitude").val(),
          "latitude": $("#latitude").val(),
          "timezone": $("#timezone").val(),
          "type": $("#tag_type").val(),
          // "country": $("#country").val(),
        },
        success: function () {
          $(".tag-add__modal").removeClass("show");
          $(".tag-modal__container").css("transition", "transform 550ms ease-in-out");
          $(".tag-modal__container").removeClass("show");
          window.location.reload()
        },
      });
    } else {
      
      $("input").each(function (i, v) {
        if ($(v).attr("required") == "required" && $(v).val() == "") {
          $(v).next().css("display", "block")
          
          if($("textarea").val() == ""){
            $("textarea").next().css("display", "block")
          }
          if ($("select").val() == "") {
            $("select").next().css("display", "block")
          }
  
        } else {
          $(v).next().css("display", "none")
  
          if($("textarea").val() != ""){
            $("textarea").next().css("display", "none")
          }
          if ($("select").val() != "") {
            $("select").next().css("display", "none")
          }
        }
      })
  
    }
  });
  
}

cancelBtn.addEventListener('click', function () {
  dashboardModal.classList.remove('show');
  modalContainer.style.transition = "transform 550ms ease-in-out";
  modalContainer.classList.remove('show');
})




function editTag(id){
  localStorage.setItem("tag_id",id)
  dashboardModal.classList.add('show');
  modalContainer.classList.add('show');
  $("#tag-save").attr("data-action", "edit-tag")
  $(".editing-img").css("display", "block")
  $("#machineID").css("pointer-events", "none").css("opacity", "0.3")

  $.get("/api/get_tag", function (result) {
    $.each(JSON.parse(result).tag, function (key, value) {
      console.log(this);
    
      $(".tag-header").html("Edit Tag")
      $(".tag-header__desc").html("Fill out the form below to edit.")
      if (this.ID == localStorage.getItem("tag_id")) {
        $("#machineID").val(this.MachineID)
        $("#name").val(this.Location.Name)
        $(`#tag_type option[value=${this.Reassociation}]`).attr('selected', 'selected');
        $("#description").val(this.Location.Description)
        $("#address").val(this.Location.Address)
        $("#longitude").val(this.Location.Longitude)
        $("#latitude").val(this.Location.Latitude)
        
        $.get("/admin/api/d/timezone/read", function(index){
          var result = JSON.parse(index).result
          $("#timezone").append(`
            <option value="">Select time zone</option>
          `)
      
          $.each(result, function(_,tz){
            $("#timezone").append(`
              <option value="`+ tz.ID +`">`+ tz.Name +`</option>
            `)
            if (value.TimeZoneID == tz.ID){
              console.log("tz id:",value.TimeZoneID,tz.ID)
              $("#timezone").val(value.TimeZoneID)
            }
          })
        })
        return false
      }
    })
  })

  $("[data-action='edit-tag']").on('click', function () {
    if ($("#name").val() != "" && $("#description").val() != "" &&
      $("#address").val() != "" && $("longitude").val() != "" &&
      $("#latitude").val() != "" && $("#timezone").val() != ""){
        
      if (this.getAttribute("data-action") == "edit-tag") {
        console.log("edit");
        $.ajax({
          url: "/api/edit_tag",
          data: {
            "id": localStorage.getItem("tag_id"),
            "name": $("#name").val(),
            "description": $("#description").val(),
            "address": $("#address").val(),
            "longitude": $("#longitude").val(),
            "latitude": $("#latitude").val(),
            "timezone": $("#timezone").val(),
            "type": $("#tag_type").val(),
          },
          success: function () {
            $(".tag-add__modal").removeClass("show");
            $(".tag-modal__container").removeClass("show");
            localStorage.removeItem("tag_id")
            window.location.reload()
          },
        });
      }
    } else {

      $("input").each(function (i, v) {
        if ($(v).attr("required") == "required" && $(v).val() == "") {

          $(v).next().css("display", "block")
          if ($("select").val() == "") {
            $("select").next().css("display", "block")
          }
        } else {
          $(v).next().css("display", "none")
          if ($("select").val() != "") {
            $("select").next().css("display", "none")
          }
        }
      })

    }
  })
}


