import React, { useCallback, useState } from "react";
import { Tabs, Tab, Box, Button } from "grommet";
import { Data } from "./data";
import styled from "styled-components";

//const Input = styled``;

const Main = () => {
  const [data, setData] = useState([
    [
      { value: "X-Axis Label", readOnly: true, width: 250 },
      { value: "Value", readOnly: true, width: 250 },
    ],
    [{ value: "Bob" }, { value: 33 }],
    [{ value: "Alice" }, { value: 14 }],
    [{ value: "Mel" }, { value: 43 }],
    [{ value: "Kip" }, { value: 50 }],
    [{ value: "Peter" }, { value: 10 }],
    [{ value: "David" }, { value: 24 }],
    [{}, {}],
    [{}, {}],
    [{}, {}],
  ]);
  const [fileName, setFileName] = useState("ShXscNXK.webm");
  const [generating, setGenerating] = useState(false);

  const generate = useCallback(async () => {
    const stringValues = data
      .filter((row) => row[1].value !== undefined && !row[1].readOnly)
      .map((row) => row[1].value);

    const step = 10;
    const values = stringValues.map((v) => parseFloat(v));
    const max = Math.max(...values);
    const maxCeil = Math.ceil(max);
    const maxCharted = (Math.floor(maxCeil / step) + 1) * step;

    setGenerating(true);
    const result = await fetch("/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ data: values, config: { max: maxCharted } }),
    });

    const fileName = await result.text();
    setGenerating(false);
    setFileName(fileName);
  }, [data]);

  return (
    <Box align="center" gap="small" margin="small">
      <Tabs alignControls="start" width="100%">
        <Tab title="Data">
          <Data data={data} setData={setData}></Data>
        </Tab>
        <Tab title="Configurations">
          <Box pad="medium">Two</Box>
        </Tab>
      </Tabs>

      <Button
        secondery
        label={generating ? "Generating..." : "Generate Video"}
        disabled={generating}
        onClick={generate}
      />

      <video controls={!generating} width="720" key={fileName} disab>
        <source src={`/videos/${fileName}`} type="video/webm" />
        Sorry, your browser doesn't support embedded webm videos. But you can
        still download it.
      </video>
    </Box>
  );
};

export { Main };
