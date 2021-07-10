import React from 'react';
import { Form, Input, Button, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';

import axios from 'axios';
import {BASE_URL} from '../constants';


function Login(props) {
    const{HandleLogin} = props
    // onfinish 是点击login button之后执行的函数
    // 处理登录逻辑
    // 前后端通信，所以用axios
    // 同时，login 之后要把login得到状态通知给<app/>
    const onFinish = (values) => {
        const{username, password} = values
        const opt = {
            method:'post',
            url:`${BASE_URL}/signin`,
            data:{
                username : username,
                password : password
            },
            headers:{"content-type" : "application/json"}

        }
        axios(opt)
        .then(response=>{
            if(response.status === 200){
                //console.log(response)
                // response里会有一个data，是后端生成的token
                const {data} = response
               // console.log(data)
                HandleLogin(data)
                message.success("Login Success!")
                //props.history.push('/home')
            }
        })
        .catch(err => {
            console.log('register falied:', err.message);
            message.error('Register Failed');

        })
        console.log('Received values of form: ', values);
    };
    return(
        <Form
           name="normal_login"
           className="login-form"
           initialValues={{
           remember: true,
      }}
      onFinish={onFinish}
    >
      <Form.Item
        name="username"
        rules={[
          {
            required: true,
            message: 'Please input your Username!',
          },
        ]}
      >
        <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Username" />
      </Form.Item>
      <Form.Item
        name="password"
        rules={[
          {
            required: true,
            message: 'Please input your Password!',
          },
        ]}
      >
        <Input
          prefix={<LockOutlined className="site-form-item-icon" />}
          type="password"
          placeholder="Password"
        />
      </Form.Item>
    

      <Form.Item>
        <Button type="primary" htmlType="submit" className="login-form-button">
          Log in
        </Button>
        <p className = "or">Or</p> 
        <a href="/register" className = 'res'>register now!</a>
      </Form.Item>
    </Form>

    );
    
     
}

export default Login;