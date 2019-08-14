import React, { Component } from 'react';
import { Pane, Menu, Heading } from 'evergreen-ui';

class SiteList extends Component {
    state = {  }
    render() { 
        return ( 
            <Pane>
                <Heading size={700} marginTop="default">
                    Sites Tracked
                </Heading>
                <Menu>
                    <Menu.Item>
                        Site #1
                    </Menu.Item>
                    <Menu.Item>
                        Site #2
                    </Menu.Item>
                </Menu>
            </Pane>
         );
    }
}
 
export default SiteList;