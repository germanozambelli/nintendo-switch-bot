import time
from random import randint

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

HOST = "127.0.0.1"  # Standard loopback interface address (localhost)
PORT = 65431  # Port to listen on (non-privileged ports are > 1023)

nx = nxbt.Nxbt()
controller_idx = nx.create_controller(nxbt.PRO_CONTROLLER)
print("Waiting switch connection")
nx.wait_for_connection(controller_idx)

print("Starting init macro")
macro_id = nx.macro(controller_idx, STARTING_MACRO)

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
                nx.macro(controller_idx, dataString)

                conn.sendall(data)
            accept(s)

accept(s)
s.close()