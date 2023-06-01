import React, {useState} from 'react';
import MDEditor from '@uiw/react-md-editor';

const Publish: React.FC = () => {

    const [value, setValue] = useState("**Hello world!!!**");

    const handleValue = (mdValue: string | undefined) => {
        console.log(mdValue)
        if (mdValue !== undefined) {
            setValue(mdValue)
        }
    }

    // https://github.com/uiwjs/react-md-editor

    return (
        <div className="container">
            <MDEditor
                value={value}
                onChange={handleValue}
            />
            <MDEditor.Markdown source={value} style={{ whiteSpace: 'pre-wrap' }} />
        </div>
    );
};

export default Publish;