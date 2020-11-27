import React from "react"

class ApiClient extends React.Component {
    static setPixel = (x: number, y: number, color: string) => {
        fetch('/api/setPixel', {
            method: 'post',
            body: JSON.stringify({'x': x, 'y': y, 'color': color})
          }).then(() => console.log("Successful")); 
    }
}


export default ApiClient;