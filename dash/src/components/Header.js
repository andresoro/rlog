import React, { Component } from 'react';
import { Pane, Button, Heading } from 'evergreen-ui';

class Header extends Component {
    render() { 
        return ( 
            <Pane display="flex" padding={16} background="tint2" borderRadius={3}>
                <Pane flex={1} alignItems="center" display="flex">
                    <Heading size={600}>ez analytics</Heading>
                 </Pane>
                 <Pane>
                    <Button marginRight={16}>Share</Button>
                    <Button appearance="primary">Export as PDF</Button>
                    </Pane>
            </Pane>
         );
    }
}
 
export default Header;