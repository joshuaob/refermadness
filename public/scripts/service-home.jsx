var ServicePanel = React.createClass({
  switchToCreate: function() {
    console.log("create from service");
    history.pushState(null, null, "/service/create");
  },
  render: function() {
    return (
      <div className="search-panel text-center">
        <div className="container">
          <SearchPage selected={this.props.service} onAddService={this.switchToCreate} />
        </div>
      </div>
    );
  }
});

var ServiceHome = React.createClass({
  render: function() {
    $(window).off("popstate").on("popstate", function() {
      window.location = window.location.href;
    });

    // strip the service name from the URL
    var serviceId = window.location.pathname.replace("/service/", "");
    var service = testData.filter(function(data) {
      return data.id === serviceId;
    })[0];

    return (
      <div className="service-home">
        <Header smallTitle={true} />
        <ServicePanel service={service} />
      </div>
    );
  }
});

React.render(
  <ServiceHome />,
  document.getElementById('content')
);