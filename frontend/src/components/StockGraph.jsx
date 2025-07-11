import React from "react";
import { Line } from "react-chartjs-2";
import { Chart as ChartJS, LineElement, PointElement, CategoryScale, LinearScale } from "chart.js";

ChartJS.register(LineElement, PointElement, CategoryScale, LinearScale);

const StockGraph = ({ stockData }) => {
  const labels = [];
  const datasets = [];

  Object.entries(stockData).forEach(([symbol, entry], i) => {
    const points = entry.history.map((pt) => ({
      x: new Date(pt.Date).toLocaleDateString(),
      y: pt.Close,
    }));

    if (labels.length === 0 && points.length > 0) {
      points.forEach((p) => labels.push(p.x));
    }

    datasets.push({
      label: symbol,
      data: points.map((p) => p.y),
      borderColor: `hsl(${(i * 37) % 360}, 70%, 50%)`,
      fill: false,
    });
  });

  const chartData = {
    labels,
    datasets,
  };

  return <Line data={chartData} />;
};

export default StockGraph;
