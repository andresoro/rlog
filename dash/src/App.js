import React, { Component } from 'react';
import Header from './components/Header';
import LineChart from './components/LineChart';
import SiteList from './components/SiteList';
import PieChart from './components/PieChart';
import RequestBox from './components/RequestBox';
import 'bootstrap-4-grid/css/grid.min.css';
import './App.css';

class App extends Component {
  render() { 
    return ( 
      <div className="bootstrap-wrapper">
        <div className="app-container container">
          <div className="row">
            <div className="col">
              <Header/>
            </div>
          </div>

          <div className="row">
            <div className="col-xs-3 col-sm-3 col-md-3 col-lg-3 col-xl-3">
              <SiteList/>
            </div>

            <div className="col-xs-9 col-sm-9 col-md-9 col-lg-9 col-xl-9">
              <div className="row">
                <div className="col-xs-6 col-sm-6 col-md-4 col-lg-4 col-xl-4">
                  <PieChart/>
                </div>

                <div className="col-xs-12 col-sm-12 col-md-6 col-lg-6 col-xl-6">
                  <LineChart/>
                </div>
              </div>

              <div className="row">
                <div className="col-xs-12 col-sm-12 col-md-12 col-lg-12 col-xl-12">
                  <RequestBox/>
                </div>
              </div>

            </div>
          </div> 
        </div>
      </div>
     );
  }
}
 
export default App;

