import react, { useEffect, useState } from 'react'
import ReactDOM from 'react-dom'


function MyComponent(){
    const[items, setItems] = useState([])

    // useEffect(() => {
    //     let data = {
    //         Name : name.value,
    //         ID : id.value,
    //         Department : dept.value
    //     };

    //     fetch('/department', {
    //         headers : {
    //             'Accept' : 'application/json',
    //             'Content-Type' : 'application/json'
    //         },
    //         method : 'POST',
    //         body : JSON.stringify(data)
    //     }).then((response) => {
    //         response.text().then((data) => {
    //             let result = JSON.parse(data)
    //             console.log(result)
    //         });
    //     }).catch((error) => {
    //         console.log(error)
    //     });
    // }, [])

    const handleSubmit = (event) => {
        const formData = new FormData(event.target);

        let jsonFormData = {};

        for (var entry of formData.entries())
        {
            jsonFormData[entry[0]] = entry[1];
        }
        jsonFormData = JSON.stringify(jsonFormData)

        // console.log(jsonFormData)

        event.preventDefault()

        // for(let [key, value] of formData.entries()){
        //     console.log(key, value)
        // }

        fetch('/department', {
                    headers : {
                        'Accept' : 'application/json',
                        'Content-Type' : 'application/json'
                    },
                    method : 'POST',
                    body : jsonFormData
                }).then((response) => {
                    response.text().then((data) => {
                        let result = JSON.parse(data)
                        console.log(result)
                    }
                    );
                }).catch((error) => {
                    // console.log(error)
                });
    }

    return(
        <div>
            <a href="/department">Department</a>

            <form onSubmit={handleSubmit}>
                <input type="text" name="name" id="name" placeholder="Enter Name" />
                <input type="number" name="id" id="id" placeholder="Enter ID" />
                <input type="text" name="department" id="department" placeholder="Enter Dept" />
                <button type="submit">Submit</button>
            </form>
        </div>
    )

}

ReactDOM.render(
    <MyComponent />,
    document.getElementById('root')
  );