var ServicePage = React.createClass({
  render: function() {
    return (
      <div className="service-area">
        <div className="view-result row">
          <div className="col-xs-12">
            <h1 className="text-center service-name">
              {this.props.data.name}
            </h1>
            <h2 className="text-center">
              <a href={this.props.data.url}>{this.props.data.url}</a>
            </h2>
            <h4 className="text-center">
              A brief description of the service.
            </h4>
          </div>
        </div>
        <div className="row random-referral-code">
          <div className="col-xs-12">
            <h3>
              Use this referral code:
            </h3>
            <h1>
              ywj-rpl
            </h1>
          </div>
        </div>
        <div className="row add-referral-code">
          <div className="col-xs-12">
            <button className="btn btn-lg btn-default">
              <span className="glyphicon glyphicon-plus" />
              Have your own code? Add it!
            </button>
          </div>
        </div>
      </div>
    );
  }
});