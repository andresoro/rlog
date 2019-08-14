import React, { Component } from 'react';
import { VictoryPie } from 'victory';
import { Heading } from 'evergreen-ui'; 

class PieChart extends Component {
    state = {  }
    render() { 
        return ( 
            <div>
                <Heading size={300} marginTop="default">
                    Top key requests
                </Heading> 
                <VictoryPie/>
            </div>
            
         );
    }
}
 
export default PieChart;