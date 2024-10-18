import matplotlib.pyplot as plt

# Function to read data from a file
def read_data(filename):
    x = []
    y = []
    with open(filename, 'r') as f:
        for row in f:
            # Split row into coordinates
            coords = row.strip().split()  # Adjust based on your delimiter
            if len(coords) == 2:  # Ensure there are exactly 2 coordinates
                try:
                    d = float(coords[0])  # Convert x coordinate to float
                    z = float(coords[1])  # Convert y coordinate to float
                    x.append(d)
                    y.append(z)
                except ValueError:
                    print(f"Invalid data encountered in file {filename}: {row.strip()}")
    return x, y

# Read data from the two text files
x1, y1 = read_data('Points.txt')
x2, y2 = read_data('Centroids.txt')

# Create the scatter plot
plt.figure(figsize=(10, 10))  # Set the figure size
plt.scatter(x1, y1, marker='+', color='blue', label='Data Set 1')
plt.scatter(x2, y2, marker='o', color='green', label='Data Set 2')

# Add labels and title
plt.title("Scatter Plot of Two Data Sets")
plt.xlabel("X-axis")
plt.ylabel("Y-axis")
plt.legend()  # Show the legend
plt.grid(True)  # Add grid for better visibility

# Show the plot
plt.show()
