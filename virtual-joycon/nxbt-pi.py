import time
from random import randint
from time import sleep
import os
import traceback

import nxbt
from nxbt import Buttons
from nxbt import Sticks

STARTING_MACRO = """
LOOP 12
    B 0.1s
    0.1s
1.5s
"""
import socket

HOST = "0.0.0.0"  # Standard loopback interface address (localhost)
PORT = 65431  # Port to listen on (non-privileged ports are > 1023)


"""Tests NXBT functionality"""
# Init
print("[1] Attempting to initialize NXBT...")
nx = None
try:
   nx = nxbt.Nxbt()
except Exception as e:
    print("Failed to initialize:")
    print(traceback.format_exc())
    exit(1)
print("Successfully initialized NXBT.\n")

# Adapter Check
print("[2] Checking for Bluetooth adapter availability...")
adapters = None
try:
    adapters = nx.get_available_adapters()
except Exception as e:
    print("Failed to check for adapters:")
    print(traceback.format_exc())
    exit(1)
if len(adapters) < 1:
    print("Unable to detect any Bluetooth adapters.")
    print("Please ensure you system has Bluetooth capability.")
    exit(1)
print(f"{len(adapters)} Bluetooth adapter(s) available.")
print("Adapters:", adapters, "\n")

# Creating a controller
print("[3] Please turn on your Switch and navigate to the 'Change Grip/Order menu.'")
input("Press Enter to continue...")

print("Creating a controller with the first Bluetooth adapter...")
cindex = None
try:
    cindex = nx.create_controller(
             nxbt.PRO_CONTROLLER,
             adapters[0])
except Exception as e:
    print("Failed to create a controller:")
    print(traceback.format_exc())
    exit(1)
print("Successfully created a controller.\n")

# Controller connection check
print("[4] Waiting for controller to connect with the Switch...")
timeout = 120
print(f"Connection timeout is {timeout} seconds for this test script.")
elapsed = 0
while nx.state[cindex]['state'] != 'connected':
    if elapsed >= timeout:
        print("Timeout reached, exiting...")
        exit(1)
    elif nx.state[cindex]['state'] == 'crashed':
        print("An error occurred while connecting:")
        print(nx.state[cindex]['errors'])
        exit(1)
    elapsed += 1
    sleep(1)
print("Successfully connected.\n")


print("Starting init macro")
macro_id = nx.macro(cindex, STARTING_MACRO)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
print("Starting Server")
s.bind((HOST, PORT))
s.listen()

def accept(s):


        conn, addr = s.accept()

        with conn:
            print(f"Connected by {addr}")
            time.sleep(2)
            while True:
                data = conn.recv(1024)
                if not data:
                    break

                dataString="".join(map(chr, data))

                print('Message received')
                print(dataString)
                nx.macro(cindex, dataString)

                conn.sendall(data)
            accept(s)

accept(s)
s.close()