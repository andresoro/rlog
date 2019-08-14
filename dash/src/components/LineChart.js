import React, { Component } from 'react';
import { VictoryLine } from 'victory';
import { Heading } from 'evergreen-ui';

class LineChart extends Component {
    state = {  }
    render() { 
        return (
            <div>
                <Heading size={300} marginTop="default">
                    Requests in past 24 hours
                </Heading> 
                <VictoryLine/>
            </div>
            
         );
    }
}
 
export default LineChart;