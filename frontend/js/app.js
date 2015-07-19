
var base_url = "http://www.jacobra.com:8003";


var CodeGrantForm = React.createClass({
  getInitialState: function() {
    return {};
  },
  handleClick: function(event) {
    var url = base_url + "/login"
    window.location.replace(url);
  },
  render: function() {
    return (
        <div>
            <button
                className="btn btn-default btn-xs"
                onClick={this.handleClick}> Code Grant
            </button>
        </div>
    );
  }
});

React.render(
  <CodeGrantForm/>,
  document.getElementById('code_grant_form')
);
