$(document).ready(function() {
    $("#zoom_pic").elevateZoom({
 //       zoomType				: "inner",
 //       cursor: "crosshair",
        tint: true,
        tintColour: '#F90',
        tintOpacity: 0.5,
        scrollZoom: true,
        gallery: 'gallery_01',
        cursor: 'pointer',
        galleryActiveClass: 'active',
        imageCrossfade: true,
        loadingIcon: '#'
    });
 
    $("#plus").click(function(){
        var tempNum=parseInt($("#num").val());
        $("#num").val(tempNum+1);
    });
    $("#minus").click(function(){
        var tempNum=parseInt($("#num").val());
        if((tempNum-1)==0){ return}
        $("#num").val(tempNum-1);
    });

    function loadXML(xmlpath) {
        var xmlDoc = null;
        if (window.ActiveXObject) {
            xmlDoc = new ActiveXObject("Microsoft.XMLDOM");
        } else if (document.implementation && document.implementation.createDocument) {
            xmlDoc = document.implementation.createDocument("", "", null);
        } else {
            alert('Your browser cannot handle this script.');
        }
        xmlDoc.async = false;
        xmlDoc.load(xmlpath);
        return xmlDoc;
    }

    $(function () {
        var xmlDoc = null;
        xmlDoc = loadXML("./home/static/js/xml/area.xml");
        var $s1 = $("#SelectProvince");
        var $s2 = $("#SelectCity");
        var $s3 = $("#SelectDistrict");
        var v1 = "省份";
        var v2 = "市";
        var v3 = "区县";
        var root = $(xmlDoc).find("address")[0];
        $(root).children("province").each(function () {
            appendOptionTo($s1, $(this).attr("name"), $(this).attr("name"), v1);
        });
        $s1.change(function () {
            $s2.html("");
            var province_node = $(root).children("province")[this.selectedIndex];
            $(province_node).children("city").each(function () {
                appendOptionTo($s2, $(this).attr("name"), $(this).attr("name"), v2);
            });
            $s2.change();
        }).change();
        $s2.change(function () {
            $s3.html("");
            var province_node = $(root).children("province")[$s1[0].selectedIndex];
            var city_node = $(province_node).children("city")[this.selectedIndex];
            $(city_node).children("country").each(function () {
                appendOptionTo($s3, $(this).attr("name"), $(this).attr("name"), v3)
            });
        }).change();
        function appendOptionTo($o, k, v, d) {
            var $opt = $("<option>").text(k).val(v);
            if (v == d) { $opt.attr("selected", "selected") }
            $opt.appendTo($o);
        }
    });


	
});
