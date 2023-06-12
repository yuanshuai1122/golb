import React, {useState} from 'react';
import MDEditor from '@uiw/react-md-editor';
import {Button, Form, Input, Modal, Space} from "antd";
import {SendOutlined} from "@ant-design/icons";

const Publish: React.FC = () => {

    const [value, setValue] = useState("**Hello world!!!**");
    const [open, setOpen] = useState(false);
    const [confirmLoading, setConfirmLoading] = useState(false);

    const showModal = () => {
        setOpen(true);
    };

    const handleOk = () => {
        setConfirmLoading(true);
        setTimeout(() => {
            setOpen(false);
            setConfirmLoading(false);
        }, 2000);
    };

    const handleCancel = () => {
        console.log('Clicked cancel button');
        setOpen(false);
    };

    const handleValue = (mdValue: string | undefined) => {
        if (mdValue !== undefined) {
            setValue(mdValue)
        }
    }

    const onFinish = (values: any) => {
        showModal()
        console.log('Success:', values);
        console.log(value)
    };

    const onFinishModal = (values: any) => {
        console.log('Success:', values);
        console.log(value)
    };

    // https://github.com/uiwjs/react-md-editor

    return (
        <>
            <Modal
                title="文章发布"
                open={open}
                onOk={handleOk}
                confirmLoading={confirmLoading}
                onCancel={handleCancel}
            >
                <Form
                    name="onPublish"
                    onFinish={onFinishModal}
                >
                    <Form.Item
                        label="文章标题"
                        name="title"
                        rules={[{ required: true, message: '请输入文章标题' }]}
                    >
                        <Input />
                    </Form.Item>
                </Form>
            </Modal>
            <Form
                name="basic"
                onFinish={onFinish}
            >
                <Form.Item>
                    <MDEditor
                        value={value}
                        onChange={handleValue}
                    />
                </Form.Item>
                <Form.Item>
                    <Button
                        icon={<SendOutlined />}
                        type="primary"
                        htmlType="submit"
                    >
                        发布
                    </Button>
                </Form.Item>
            </Form>

        </>

    );
};

export default Publish;