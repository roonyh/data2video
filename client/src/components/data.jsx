import React from "react";
import { Box } from "grommet";
import ReactDataSheet from "react-datasheet";

const Data = ({ data, setData }) => {
  return (
    <Box>
      <ReactDataSheet
        data={data}
        valueRenderer={(cell) => cell.value}
        onCellsChanged={(changes) => {
          const newData = [...data];
          changes.forEach(({ row, col, value }) => {
            newData[row] = [...newData[row]];
            newData[row][col] = { ...newData[row][col], value };
          });
          setData(newData);
        }}
      />
    </Box>
  );
};

export { Data };
