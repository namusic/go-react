import React, {Component} from 'react';
import ChannelForm from './ChannelForm.jsx';
import ChannelList from './ChannelList.jsx';
import PropTypes from 'prop-types';

class ChannelSection extends Component{
  render(){
    return(
      <div className='support panel panel-primary'>
        <div className='panel-heading'>
          <strong>Channel</strong>
        </div>
        <div>
          <ChannelList {...this.props} />
          <ChannelForm {...this.props} />
        </div>
      </div>
    )
  }

}

ChannelSection.propTypes = {
  channels: PropTypes.array.isRequired,
  setChannel: PropTypes.func.isRequired,
  addChannel: PropTypes.func.isRequired,
  activeChannel: PropTypes.object.isRequired
}

export default ChannelSection