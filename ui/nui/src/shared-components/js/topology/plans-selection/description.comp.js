import React from 'react';


var Description = React.createClass({
  render(){
    return(
      <div className={`${this.props.className}__description`} style={{ height: this.props.plansCount*42 }}>
        <span>{this.props.text}</span>
      </div>
    )
  }
});

Description.propTypes = {
  text: React.PropTypes.string.isRequired,
  className: React.PropTypes.string.isRequired
}

export default Description;
