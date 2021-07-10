import React from 'react';

import Login from './Login.js';
import Register from './Register.js';
import Home from './Home.js'; 

import {Route, Switch, Redirect} from 'react-router-dom';

// 只用Route会有问题：
// 当你的path为以下形式：path= /:register, ：代表参数，则当我们输入/home的路径的时候会出现register和home两个页面
// 而我们想要的只是其中之一
// 用Switch可以解决这个问题， 精确匹配
// route 从上打下去匹配，找到最先匹配的停止


function Main(props) {
    const { isLogin, HandleLogin } = props

    const showLogin = () =>{
        // 通过isLogin判断是否login， show home component
        return isLogin?
        <Redirect to='/home'/> 
        :
        <Login HandleLogin = {HandleLogin}/>

    }
    return (
        <div>
            <Switch>
                {/* 以下方式传入的是组件名，不是<Login/> 这样的组件本身，所以不能给子组件传props，这时候要用到render
                    render 可以接受一个回调函数 */}
                <Route path='/register' component = {Register}/>
                <Route exact path='/login' render = {showLogin}/>
                <Route exact path = '/home' component = {Home}/>
            </Switch>
        </div>
        
        
      
        

    );
}

export default Main;