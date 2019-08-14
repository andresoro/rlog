import React, { Component } from 'react';
import { VictoryPie, VictoryLine } from 'victory';
import 'bootstrap-4-grid/css/grid.min.css';
import './App.css';

class App extends Component {
  render() { 
    return ( 
      <div className="bootstrap-wrapper">
        <div className="app-container container">
          <div className="row">
            <div className="col-xs-6 col-sm-6 col-md-6 col-lg-6 col-xl-6">
              <h1> EZ Analytics</h1>
            </div>

            <div className="col-xs-6 col-sm-6 col-md-6 col-lg-6 col-xl-6">
              <button> Share </button>
              <button> Export to PDF</button>
            </div>
          </div>

          <div className="row">
            <div className="col-xs-3 col-sm-3 col-md-3 col-lg-3 col-xl-3">
              <h2>Site List Component</h2>
            </div>

            <div className="col-xs-9 col-sm-9 col-md-9 col-lg-9 col-xl-9">
              <div className="row">
                <div className="col-xs-6 col-sm-6 col-md-4 col-lg-4 col-xl-4">
                  <VictoryPie/>
                </div>

                <div className="col-xs-12 col-sm-12 col-md-6 col-lg-6 col-xl-6">
                  <VictoryLine/>
                </div>
              </div>

              <div className="row">
                <div className="col-xs-12 col-sm-12 col-md-12 col-lg-12 col-xl-12">
                  <h2> Latest requests </h2>
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

