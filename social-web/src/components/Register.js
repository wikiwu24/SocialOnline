import React from 'react';
import { Form, Input, Button,message } from 'antd';

import axios from 'axios';

import { BASE_URL } from '../constants';


const formItemLayout = {
    labelCol: {
      xs: {
        span: 24,
      },
      sm: {
        span: 8,
      },
    },
    wrapperCol: {
      xs: {
        span: 24,
      },
      sm: {
        span: 16,
      },
    },
  };
  const tailFormItemLayout = {
    wrapperCol: {
      xs: {
        span: 24,
        offset: 0,
      },
      sm: {
        span: 16,
        offset: 8,
      },
    },
  };

function Register(props) {
    const [form] = Form.useForm();
    
    // send data to the server
    const onFinish = (values) => {
        // get username/password
        const{ username, password} = values;
        // send data to server : axios
        const opt = {
            method :'post',
            url:`${BASE_URL}/signup`,
            data:{
                username : username,
                password : password
            },
            headers:{"content-type" : "application/json"}
        }
        // axios接受opt作为参数，then返回成功的回调，catch返回失败回调
        axios(opt)
        .then(response => {
            console.log(response);
            if(response.status === 200){
                console.log("register success!");
                message.success('Register Successed!')
                // 注册成功后还有个问题： 要回退到login页面进行注册
                // 对于Route的页面，有一个props叫history，可以push要退回的页面到栈顶，然后我们就能够退回了
                // 页面的地址就是我们定义的route的path
                // 没有绑定route就没有history props
                props.history.push('/login')
            }
        })
        .catch(err => {
            console.log('register falied:', err.message);
            message.error('Register Failed');
        })


        console.log('Received values of form: ', values);
    };

    return (
        <Form
        {...formItemLayout}
        form={form}
        className = "register"
        name="register"
        onFinish={onFinish}
        initialValues={{
          residence: ['zhejiang', 'hangzhou', 'xihu'],
          prefix: '86',
        }}
        scrollToFirstError
      >
        <Form.Item
          name="username"
          label="Username"
          rules={[
            {
              required: true,
              message: 'Please input your Username!',
            },
          ]}
        >
          <Input />
        </Form.Item>
  
        <Form.Item
          name="password"
          label="Password"
          rules={[
            {
              required: true,
              message: 'Please input your password!',
            },
          ]}
          hasFeedback
        >
          <Input.Password />
        </Form.Item>
  
        <Form.Item
          name="confirm"
          label="Confirm Password"
          dependencies={['password']}
          hasFeedback
          rules={[
            {
              required: true,
              message: 'Please confirm your password!',
            },
            ({ getFieldValue }) => ({
              validator(_, value) {
                if (!value || getFieldValue('password') === value) {
                  return Promise.resolve();
                }
  
                return Promise.reject(new Error('The two passwords that you entered do not match!'));
              },
            }),
          ]}
        >
          <Input.Password />
        </Form.Item>
        <Form.Item {...tailFormItemLayout}>
        <Button type="primary" htmlType="submit">
          Register
        </Button>
      </Form.Item>
    </Form>
    );
}

export default Register;