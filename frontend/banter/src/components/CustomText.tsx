import React, { useEffect } from "react";
import { useState } from "react";

function CustomText() {
  const [records, setRecords] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/")
      .then((response) => response.json())
      .then((data) => setRecords(data))
      .catch((err) => console.log(err));
  }, []);

  return (
    <div>
      <h1>{records}</h1>
    </div>
  );
}

export default CustomText;
