
import React, {useState} from 'react';
import HeadBar from './HeadBar.js';
import {TOKEN_KEY} from '../constants';
import Main from './Main.js';


function App() {
  // useState 可以让function component使用state，
  // 该函数返回： 一个array， 第一个是: state 的初始值（可在useState中传入），
  // 第二个是一个function，是该状态的setState function, 用于修改该状态值
  // 一个状态对应一个useState

 // 定义一个是否登录的状态
 // 有token就是登录，所以isLogin的初始值是判断是否有token
 // 在login， logout funcition 中 用setLogin function对isLogin状态进行相应的变化
  const [isLogin, setLogin] = useState(
    localStorage.getItem(TOKEN_KEY)? true : false
  );

 const loggout = ()=>{
   // 在logout 后将localstorage 中的token删除
   localStorage.removeItem(TOKEN_KEY);
   setLogin(false);
 }

 const loggedIn = token =>{
   if(token) {
     localStorage.setItem(TOKEN_KEY, token);
     setLogin(true);
   }
 }

 // 如何更新islogin的变化？是组件的状态值，在function component中没有state的概念，所以要用到hooks


  return (
    <div className="App">
      <HeadBar 
        isLogin={isLogin} 
        HandleLogout = {loggout} 
      />
      <Main
        isLogin = {isLogin}
        HandleLogin = {loggedIn}
      />
    </div>
  );
}

export default App;
