$(document).ready(function() {
    var editaccount=$("#editaccount");
    editaccount.click(function () {
        $("#accoutcontent").load("index.html?action=EditAccount");
    });

    var updatepassword=$("#updatepassword");
    updatepassword.click(function () {
        $("#accoutcontent").load("index.html?action=UpdatePassword");
    });

    var accountlevel=$("#accountlevel");
    accountlevel.click(function () {
        $("#accoutcontent").load("index.html?action=accountlevel");
    });

    var manageaddress=$("#manageaddress");
    manageaddress.click(function () {
        $("#accoutcontent").load("index.html?action=manageaddress");
    });

    var loveproduction=$("#loveproduction");
    loveproduction.click(function () {
        $("#accoutcontent").load("index.html?action=loveproduction");
    });
    var ordermanage=$("#ordermanage");
    ordermanage.click(function () {
        $("#accoutcontent").load("index.html?action=ordermanage");
    });
});
