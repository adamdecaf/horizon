// analyze app.js
(function () {

  function GET(url, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", url, true);
    xhr.onload = function (e) {
      if (xhr.readyState === 4) {
        if (xhr.status === 200) {
          var text = xhr.responseText;
          callback(text);
          // console.log(text);
        } else {
          console.error(xhr.statusText);
        }
      }
    };
    xhr.onerror = function (e) {
      console.error(xhr.statusText);
    };
    xhr.send(null);
  }

  // cached bindings
  var cities_binding;

  // Search Cities
  window.search_cities = function() {
    var query = document.querySelector("#city");
    GET("/cities?q=" + encodeURIComponent(query.value),
        function(res) {
          if (cities_binding) {
            cities_binding.update(JSON.parse(res));
          } else {
            var container = document.querySelector("#cities-search");
            container.style.display = 'block';
            cities_binding = rivets.bind(document.querySelector("#cities-list"), JSON.parse(res));
          }
        });
  };

})();
