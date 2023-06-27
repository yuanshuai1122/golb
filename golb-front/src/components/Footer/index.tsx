import React from 'react';
import Typography from "@mui/material/Typography";

const Footer = () => {
    return (
        <div style={{display: "flex",flexDirection: "column",alignItems: "center",justifyContent:"flex-end",flexGrow: 1, marginTop: 20, marginBottom: 20}}>
            <footer id="footer">
                <Typography variant="subtitle1" sx={{textAlign: 'center'}} gutterBottom>
                    © 2023 亭子 冀ICP备xxxxxxxxx号
                </Typography>
                <Typography variant="subtitle2" sx={{textAlign: 'center'}} gutterBottom>
                    Designed by YuanShuai
                </Typography>
            </footer>
        </div>
    );
};

export default Footer;