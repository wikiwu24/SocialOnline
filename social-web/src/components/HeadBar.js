import React from 'react';
import logo from '../assets/images/logo.svg';
import {LogoutOutlined} from '@ant-design/icons';

function HeadBar(props){
  // props 会传给我islogin， 来决定是否显示logout button

  const{isLogin, HandleLogout} = props;

    return(
        <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <span className = "App-title">Social Web</span>
        {
          isLogin? 
            <span className = "logout-bar">
              <span className = "logout-title">Log out</span>   
              <LogoutOutlined className = "logout"
                              onClick = {HandleLogout}
                             />
            </span>
             :null
        }
        
      </header>

    )
    
    

}
export default HeadBar;