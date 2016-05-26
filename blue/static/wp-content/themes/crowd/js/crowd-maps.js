function createMap(mapid,mapdata) {
    // Create an array of styles.
    var styles =
    [
        {
            stylers: [
                { saturation: -100 }
            ]
        }
    ];

    // Create a new StyledMapType object, passing it the array of styles,
    // as well as the name to be displayed on the map type control.
    var styledMap = new google.maps.StyledMapType(styles,{name: "Map"});


    var mapOptions = {
        center: mapdata.center,
        zoom: mapdata.zoom,
        mapTypeControlOptions: {
            mapTypeIds:['map_style']
        }
    };
    mapdata.map = new google.maps.Map(document.getElementById(mapid),mapOptions);

    mapdata.map.mapTypes.set('map_style',styledMap);
    mapdata.map.setMapTypeId('map_style');

    var marker = new google.maps.Marker({
        position: mapdata.center,
        map: mapdata.map,
        title:mapdata.title
    });


    google.maps.event.addDomListener(window, "resize", function() {
        google.maps.event.trigger(mapdata.map, "resize");
        mapdata.map.setCenter(mapdata.center); 
    });
}


function initialize() {
    var mapid;
    for (mapid in maps) {
        if(maps.hasOwnProperty(mapid)) {
            createMap(mapid, maps[mapid]);
        }
    }
}

var maps = {
    'uk-map': { map: null, zoom: 14, center: new google.maps.LatLng( 50.74101, -1.89888 ), title: "Crowd United Kingdom" },
    'sf-map': { map: null, zoom: 14,center: new google.maps.LatLng( 37.784072, -122.394645 ), title: "Crowd San Francisco" },
    'db-map': { map: null, zoom: 14,center: new google.maps.LatLng( 25.191882, 55.284418 ), title: "Crowd Dubai" }
}

google.maps.event.addDomListener(window, 'load', initialize);


